package main

import (
	"context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/bayu-aditya/myfacilities-backend/graph"
	"github.com/bayu-aditya/myfacilities-backend/graph/generated"
	"github.com/bayu-aditya/myfacilities-backend/lib/middleware"
	"github.com/bayu-aditya/myfacilities-backend/lib/service"
	"github.com/bayu-aditya/myfacilities-backend/lib/variable"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func mainHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	godotenv.Load("./devel.env")
	variable.InitializationVariableEnvironment()

	mongoClient := service.InitializationMongo(ctx)
	defer mongoClient.Disconnect(ctx)

	router := gin.Default()
	router.Use(middleware.GinContextToMiddleware())

	router.POST("/query", mainHandler())

	router.Run(":8080")
}

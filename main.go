package main

import (
	"context"
	"fmt"

	"github.com/bayu-aditya/myfacilities-backend/lib/service"
	"github.com/bayu-aditya/myfacilities-backend/lib/variable"

	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	godotenv.Load("./devel.env")
	variable.InitializationVariableEnvironment()

	mongoClient := service.InitializationMongo(ctx)
	defer mongoClient.Disconnect(ctx)

	fmt.Println("success")
}

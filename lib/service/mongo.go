package service

import (
	"context"
	"log"

	"github.com/bayu-aditya/myfacilities-backend/lib/variable"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB address for connect to collection
var MongoDB *mongo.Database

// InitializationMongo for first time connection
func InitializationMongo(ctx context.Context) *mongo.Client {
	conf := variable.Mongo
	client, err := mongo.NewClient(options.Client().ApplyURI(conf.URI))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	MongoDB = client.Database(conf.DB)
	return client
}

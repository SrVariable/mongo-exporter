package main

import (
	"context"
	"fmt"
	"os"

	router "github.com/SrVariable/mongo-exporter/api"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const defaultMongoURI = "mongodb://localhost:27017"

func connectDatabase() (*mongo.Client, context.Context, context.CancelFunc, error) {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = defaultMongoURI
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func closeConnection(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func main() {
	client, c, cancel, err := connectDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer closeConnection(client, c, cancel)

	router.Run(client, ":8080")
}

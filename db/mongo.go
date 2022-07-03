package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"rumah-hewan/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MgoConn ...
func MgoConn() *mongo.Client {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("config tidak di temukan", err)
	}

	var (
		client   *mongo.Client
		mongoURL = config.Database.Mongo.URI
	)

	// Initialize a new mongo client with options
	client, err = mongo.NewClient(options.Client().ApplyURI(mongoURL))
	if err != nil {
		log.Fatalf("Error when instantiate mongodb connection: %s", err.Error())
	}

	// Connect the mongo client to the MongoDB server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	defer cancel()
	if err != nil {
		panic(err)
	}

	// Ping MongoDB
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		panic(fmt.Sprintf("could not ping to mongo db service: %v\n", err))
	}

	return client
}

// MgoCollection call a collection with passing client value
func MgoCollection(coll string, client *mongo.Client) *mongo.Collection {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("config tidak di temukan", err)
	}

	return client.Database(config.Database.Mongo.Db).Collection(coll)
}

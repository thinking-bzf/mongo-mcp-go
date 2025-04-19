package client

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"mcp/app/configs"
)

var (
	DB          *mongo.Database
	MongoClient *mongo.Client
)

// ConnectMongo connects to a MongoDB database using the provided configuration.
func ConnectMongo(config configs.MongoConfig) {
	host := config.Host
	port := config.Port
	user := config.User
	password := config.Password
	database := config.Database

	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s:%d", user, password, host, port),
	)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("failed to ping to mongodb: %v", err)
	}
	log.Println("Connected to MongoDB!")

	// if the database does not exist, MongoDB create it
	DB = client.Database(database)
	MongoClient = client
}

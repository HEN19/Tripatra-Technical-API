package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func Connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // MongoDB connection URI
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Assign to global MongoClient for reuse
	MongoClient = client
	log.Println("Connected to MongoDB!")
	return client
}

func GetMongoCollection(dbName, collectionName string) *mongo.Collection {
	if MongoClient == nil {
		log.Fatal("MongoDB client is not initialized. Call Connect() first.")
	}
	return MongoClient.Database(dbName).Collection(collectionName)
}

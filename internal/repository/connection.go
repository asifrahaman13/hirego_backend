package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository[T any] struct {
	db *mongo.Client
}

// This repository uses sql server
func InitializeDB() (*mongo.Client, error) {

	// Load the .env file.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file is specified.")
	}

	// Fetch the uri from the environment variable.
	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		return nil, errors.New("you must set your 'MONGODB_URI' environment variable")
	}

	// Connect to the mongodb database.
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	// Error handling.
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %s", err)
	}

	fmt.Println("Connected to MongoDB!")
	// Return either the client or nil in case of error.
	return client, nil
}
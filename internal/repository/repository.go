package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/asifrahaman13/hirego/internal/core/domain"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
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

// Function to store the userdata
func (r *repository[T]) Create(model T) (string, error) {
	coll := r.db.Database("hirego").Collection("users")
	_, err := coll.InsertOne(context.TODO(), model)

	if err != nil {
		return "", err
	}

	return "User signed up successfully", nil
}

func (r *repository[T]) GetByEmail(username string) (interface{}, error) {
	coll := r.db.Database("hirego").Collection("userprofile")

	filter := bson.D{{Key: "email", Value: username}}

	var userInformation domain.UserInformation

	err := coll.FindOne(context.TODO(), filter).Decode(&userInformation)
	if err != nil {
		return nil, err
	}

	return userInformation, nil
}

func (r *repository[T]) InsertData(username string, workinforamtion interface{}, collection string) (string, error) {
	coll := r.db.Database("hirego").Collection(collection)
	_, err := coll.InsertOne(context.TODO(), workinforamtion)

	if err != nil {
		return "", err
	}

	return "Work information added successfully", nil
}

func (r *repository[T]) GetData(username string, collection string) (interface{}, error) {
	coll := r.db.Database("hirego").Collection(collection)

	filter := bson.D{{Key: "useremail", Value: username}}

	var result domain.WorkInformation
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

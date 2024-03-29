package repository

import (
	"context"
	"fmt"

	"github.com/asifrahaman13/hirego/internal/core/domain"
	"github.com/asifrahaman13/hirego/internal/helper"
	"go.mongodb.org/mongo-driver/bson"
	"errors"
	"github.com/joho/godotenv"
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

func (r *repository[T]) GetData() ([]domain.User, error) {
	coll := r.db.Database("hirego").Collection("users")
	filter := bson.D{}

	cursor, err := coll.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results []domain.User
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	for _, user := range results {
		fmt.Printf("User Email: %s, \n", user.Email)
	}

	return results, nil
}

// Function to store the userdata
func (r *UserRepository) SignUp(user domain.User) (string, error) {
	coll := r.db.Database("hirego").Collection("users")
	_, err := coll.InsertOne(context.TODO(), user)

	if err != nil {
		return "", err
	}

	return "User signed up successfully", nil
}

func (r *UserRepository) Login(user domain.User) (domain.AccessToken, error) {

	// Call the create token function to generate the access token.
	access_token, err := helper.CreateToken(user.Email)

	// Return the access token.
	if err != nil {
		return domain.AccessToken{}, err
	}

	// Return the access token.
	return domain.AccessToken{Token: access_token}, nil
}

func (r *UserRepository) ProtectedRoute(token string) (string, error) {
	claims, err := helper.VerifyToken(token)
	if err != nil {
		return "", err
	}

	return claims["username"].(string), nil
}

func (r *UserRepository) UserInformation(user domain.UserInformation) (string, error) {
	// Define the collection.
	coll := r.db.Database("hirego").Collection("userinformation")

	fmt.Println(user)
	// Insert the document into the collection.
	_, err := coll.InsertOne(context.TODO(), user)

	if err != nil {
		return "", err
	}

	return "User information stored successfully", nil
}

func (r *UserRepository) GetUserInformation(email string) (domain.UserInformation, error) {
	// Define the collection.
	coll := r.db.Database("hirego").Collection("userinformation")

	var user domain.UserInformation

	// Ordered representation of a BSON filter. Find all the documents with the email.
	filter := bson.D{{Key: "email", Value: email}}

	// Find the document with the email.
	err := coll.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		return domain.UserInformation{}, err
	}

	return user, nil
}

func (r *UserRepository) SetUserWrorkInformation(username string, workinformation domain.WorkInformation) (string, error) {

	workinformation.Useremail = username

	// Define the collection.
	coll := r.db.Database("hirego").Collection("workinformation")

	_, err := coll.InsertOne(context.TODO(), workinformation)

	if err != nil {
		return "Something went wrong", err
	}

	return "Data stored successfully", nil
}

func (r *UserRepository) GetUserWorkInformation(domain.UserName) (domain.WorkInformation, error) {
	// Define the collection.
	coll := r.db.Database("hirego").Collection("workinformation")

	var workinformation domain.WorkInformation

	// Find the document with the email.
	err := coll.FindOne(context.TODO(), bson.D{}).Decode(&workinformation)

	if err != nil {
		return domain.WorkInformation{}, err
	}

	return workinformation, nil
}

package repository

import (
	"context"
	"fmt"
	"github.com/asifrahaman13/hirego/internal/core/domain"
	"github.com/asifrahaman13/hirego/internal/helper"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserRepo *UserRepository

type UserRepository struct {
	*repository[domain.User]
}

// BaseRepository implements ports.UserRepository.
func (r *UserRepository) BaseRepository(*domain.User) {
	panic("unimplemented")
}

func (r *UserRepository) Initialize(db *mongo.Client) *UserRepository {
	UserRepo = &UserRepository{
		repository: &repository[domain.User]{db: db},
	}

	return UserRepo
}

// Function to store the userdata
func (r *UserRepository) SignUp(user *domain.User) (string, error) {
	coll := r.db.Database("hirego").Collection("users")
	_, err := coll.InsertOne(context.TODO(), user)

	if err != nil {
		return "", err
	}

	return "User signed up successfully", nil
}

func (r *UserRepository) Login(user *domain.User) (*domain.AccessToken, error) {

	// Call the create token function to generate the access token.
	access_token, err := helper.CreateToken(user.Email)

	// Return the access token.
	if err != nil {
		return nil, err
	}

	// Return the access token.
	return &domain.AccessToken{Token: access_token}, nil
}

func (r *UserRepository) ProtectedRoute(token string) (string, error) {
	claims, err := helper.VerifyToken(token)
	if err != nil {
		return "", err
	}

	return claims["username"].(string), nil
}

func (r *UserRepository) UserInformation(user *domain.UserInformation) (string, error) {
	coll := r.db.Database("hirego").Collection("userinformation")

	fmt.Println(user)
	_, err := coll.InsertOne(context.TODO(), user)

	if err != nil {
		return "", err
	}

	return "User information stored successfully", nil
}

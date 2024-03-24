package infrastructure

import (
	"context"
	"fmt"
	"github.com/asifrahaman13/hirego/src/internal/domain"
	"github.com/asifrahaman13/hirego/src/internal/services"
)

type AuthRepository struct {
	// some fields
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (r *AuthRepository) Signup(user *domain.User) (interface{}, error) {
	mongodbClient, err := services.Services()

	if err != nil {
		return nil, err
	}

	coll := mongodbClient.Database("hirego").Collection("users")

	result, err := coll.InsertOne(context.TODO(), user)

	fmt.Println("result: ", result)

	if err != nil {
		return nil, err
	}

	userMap := map[string]interface{}{
		"firstname": user.FirstName,
	}

	return userMap, nil
}

package repository

import (
	"context"
	"fmt"
	"github.com/asifrahaman13/hirego/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository[T]) GetData() ([]*domain.User, error) {
	coll := r.db.Database("hirego").Collection("users")
	filter := bson.D{}

	cursor, err := coll.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var results []*domain.User
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	for _, user := range results {
		fmt.Printf("User ID: %s, Name: %s\n", user.Email, user.FirstName)
	}

	return results, nil
}

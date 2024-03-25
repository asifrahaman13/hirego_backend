package repository

import (
	"github.com/asifrahaman13/hirego/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository[T any] struct {
	// db *mongo.Client
}

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
		repository: &repository[domain.User]{},
	}

	return UserRepo
}

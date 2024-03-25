package repository

import (
	"github.com/asifrahaman13/hirego/internal/core/domain"
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

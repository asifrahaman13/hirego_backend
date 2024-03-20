// infrastructure/user_repository.go
package infrastructure

import (
	"encoding/json"

	"github.com/asifrahaman13/clean/src/internal/domain"
)

type UserRepository struct {
	// Any dependencies or configuration needed
}

// Save implements domain.UserRepository.
func (r *UserRepository) Save(user *domain.User) error {
	panic("unimplemented")
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindByID(id int) (*domain.User, error) {
	// Simulating a database query to find user by ID
	// For demonstration purposes, we'll just return a hardcoded user

	// Dummy user data
	user := &domain.User{
		ID:   1,
		Name: "John Doe",
		Age:  30,
	}

	// Serialize user object to JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	// Deserialize JSON back to user object
	var newUser domain.User
	if err := json.Unmarshal(userJSON, &newUser); err != nil {
		return nil, err
	}

	return &newUser, nil
}

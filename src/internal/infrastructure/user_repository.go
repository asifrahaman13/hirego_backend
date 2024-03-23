// infrastructure/user_repository.go
package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/asifrahaman13/hirego/src/internal/domain"
)

type UserRepository struct {
	// Any dependencies or configuration needed
}

// FindAll implements domain.UserRepository.
// FindAll implements domain.UserRepository.
func (r *UserRepository) FindAll() ([]*domain.User, error) {
	user := &domain.User{
		ID:   1,
		Name: "John Doe sample data here.",
		Age:  30,
	}

	// Marshal the user object to JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	// Print the JSON string (optional)
	fmt.Println(string(userJSON))

	// Return a slice of pointers to user objects
	return []*domain.User{user}, nil
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

	fmt.Print(*user)


	return user, nil
}


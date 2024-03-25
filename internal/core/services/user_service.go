package service

import (
	"github.com/asifrahaman13/hirego/internal/core/domain"
	"github.com/asifrahaman13/hirego/internal/core/ports"
)

type userService struct {
	repo ports.UserRepository
}

func InitializeUserService(r ports.UserRepository) *userService {
	return &userService{
		repo: r,
	}
}

func (s *userService) GetAllUsers() ([]*domain.User, error) {

	data, err := s.repo.GetData()

	if err != nil {
		panic(err)
	}

	return data, nil

}

// Service to signup a user.
func (s *userService) Signup(user *domain.User) (string, error) {

	// Call the sign up repo to insert the data of the user.
	message, err := s.repo.SignUp(user)

	if err != nil {
		panic(err)
	}

	// Return the success message.
	return message, nil
}


func (s *userService) Login(user *domain.User) (*domain.AccessToken, error) {
	
	// Call the login repo to insert the data of the user.
	token, err := s.repo.Login(user)

	if err != nil {
		panic(err)
	}

	// Return the success message.
	return token, nil
}
package service

import (
	"fmt"

	"github.com/asifrahaman13/hirego/internal/core/domain"
	"github.com/asifrahaman13/hirego/internal/core/ports"
	"github.com/asifrahaman13/hirego/internal/helper"
)

type userService struct {
	repo ports.UserRepository
}

func InitializeUserService(r ports.UserRepository) *userService {
	return &userService{
		repo: r,
	}
}

// func (s *userService) Signup(user domain.User) (string, error) {
// 	return s.repo.Create(user)
// }

// Service to signup a user.
func (s *userService) Signup(user domain.User) (string, error) {

	// Call the sign up repo to insert the data of the user.
	message, err := s.repo.Create(user)

	if err != nil {
		panic(err)
	}

	// Return the success message.
	return message, nil
}

func (s *userService) Login(user domain.User) (domain.AccessToken, error) {

	// Call the login repo to insert the data of the user.
	token, err := helper.CreateToken(user.Email)

	if err != nil {
		panic(err)
	}

	// Return the success message.

	accessToken := domain.AccessToken{
		Token: token,
	}
	return accessToken, nil
}

func (s *userService) SetUserWrorkInformation(email string, workInformation domain.WorkInformation) (string, error) {

	// Call the login repo to insert the data of the user.
	message, err := s.repo.InsertData(email, workInformation, "workinformation")

	if err != nil {
		panic(err)
	}

	// Return the success message.
	return message, nil
}


func (s *userService) GetUserWorkInformation(username string) (domain.WorkInformation, error) {

	
	// Call the login repo to insert the data of the user.
	workInformation, err := s.repo.GetData(username, "workinformation")

	if err != nil {
		panic(err)
	}

	// Perform a type assertion to convert workInformation to domain.WorkInformation.
	info, ok := workInformation.(domain.WorkInformation)
	if !ok {
		// Handle the case where the type assertion fails.
		return domain.WorkInformation{}, fmt.Errorf("failed to convert workInformation to domain.WorkInformation")
	}

	// Return the success message.
	return info, nil
}

func (s *userService) GetProfileInformation(username string) (domain.UserInformation, error) {

	// Call the login repo to insert the data of the user.
	userInformation, err := s.repo.GetByEmail(username)

	if err != nil {
		panic(err)
	}

	// Perform a type assertion to convert workInformation to domain.WorkInformation.
	info, ok := userInformation.(domain.UserInformation)
	if !ok {
		// Handle the case where the type assertion fails.
		return domain.UserInformation{}, fmt.Errorf("failed to convert workInformation to domain.WorkInformation")
	}

	// Return the success message.
	return info, nil
}

func (s *userService) SetUserProfileInformation(email string, userInformation domain.UserInformation) (string, error) {

	// Call the login repo to insert the data of the user.
	message, err := s.repo.InsertData(email, userInformation, "userprofile")

	if err != nil {
		panic(err)
	}

	// Return the success message.
	return message, nil
}
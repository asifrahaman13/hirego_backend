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

/*
The setUserWorkInformation function is used to set the work information of the user.
This will be public information that will be visible to all the hr managers who signs up in the platform.
*/
func (s *userService) SetUserWrorkInformation(email string, workInformation domain.WorkInformation) (string, error) {

	workInformation.Useremail = email

	// Call the login repo to insert the data of the user.
	message, err := s.repo.InsertData(workInformation, "workinformation")

	if err != nil {
		panic(err)
	}

	if !message {
		return "Some error occured", nil
	}

	// Return the success message.
	return "User work information stored successfully", nil
}

/*
The getUserWorkInformation function is used to get the work information of the user.
This will mainly be used by the HR managers to view the work information of the user.
*/
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

/*
The GetProfileInformation function is used to get the profile information of the user.
This will be completely private and only the ueer will be able to view this information. This is one to one.
*/
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

/*
This function is used to set the profile information of the user. Users will set them for profile in our website.
*/
func (s *userService) SetUserProfileInformation(email string, userInformation domain.UserInformation) (string, error) {

	userInformation.Email = email

	// Call the login repo to insert the data of the user.
	message, err := s.repo.InsertData(userInformation, "userprofile")

	if err != nil {
		panic(err)
	}

	if !message {
		return "Some error occurred", nil
	}

	// Return the success message.
	return "Your profile information stored successfully", nil
}

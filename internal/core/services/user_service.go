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
    
	// Call the get data repo to get the data of the user.
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

func (s *userService) ProtectedRoute(token string) (string, error) {

	// Call the login repo to insert the data of the user.
	message, err := s.repo.ProtectedRoute(token)

	if err != nil {
		panic(err)
	}

	// Return the success message.
	return message, nil
}

func (s *userService) UserInformation(user *domain.UserInformation) (string, error) {

	// Call the login repo to insert the data of the user.
	message, err := s.repo.UserInformation(user)

	if err != nil {
		panic(err)
	}

	// Return the success message.
	return message, nil
}

func (s *userService) GetUserInformation(email string) (*domain.UserInformation, error) {

	// Call the login repo to insert the data of the user.
	user, err := s.repo.GetUserInformation(email)

	if err != nil {
		panic(err)
	}

	// Return the success message.
	return user, nil
}

func (s *userService) SetUserWrorkInformation(username string, workinformation *domain.WorkInformation) (string, error) {

	// Call the login repo to insert the data of the user.
	_, err := s.repo.SetUserWrorkInformation(username, workinformation)

	if err != nil {
		panic(err)
	}

	// Return the success message.
	return "Data is stored successfully", nil
}


func (s *userService) GetUserWorkInformation(username *domain.UserName) (*domain.WorkInformation, error) {
	
	// Call the login repo to insert the data of the user.
	workinformation, err := s.repo.GetUserWorkInformation(username)

	if err != nil {
		panic(err)
	}

	// Return the success message.
	return workinformation, nil
}
package service

import (
	"fmt"
	"github.com/asifrahaman13/hirego/internal/core/ports"
)

type userService struct {
	repo ports.UserRepository
}

func InitializeUserService(r ports.UserRepository) *userService{
	return &userService{
		repo:r,
	}
}

func (s *userService) GetAllUsers() (map[string]interface{}, error) {
	fmt.Print("called")

	data := map[string]interface{}{
		"name": "hello",
	}

	return data, nil

}

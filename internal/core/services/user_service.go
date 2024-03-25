package service

import (
	"fmt"
	"github.com/asifrahaman13/hirego/internal/core/ports"
	"github.com/asifrahaman13/hirego/internal/core/domain"
)

type userService struct {
	repo ports.UserRepository
}

func InitializeUserService(r ports.UserRepository) *userService{
	return &userService{
		repo:r,
	}
}

func (s *userService) GetAllUsers()  ([]*domain.User, error) {
	fmt.Print("called")


	data, err:=s.repo.GetData()

	if err!=nil{
		panic(err)
	}

	return data, nil

}

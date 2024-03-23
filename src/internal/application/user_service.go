// application/user_service.go
package application

import "github.com/asifrahaman13/clean/src/internal/domain"

type UserService struct {
    UserRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) *UserService {
    return &UserService{UserRepository: userRepository}
}

func (s *UserService) GetUserByID(id int) (*domain.User, error) {
    return s.UserRepository.FindByID(id)
}

func (s *UserService) SaveUser(user *domain.User) error {
    return s.UserRepository.Save(user)
}

func (s *UserService) GetUsers() ([]*domain.User, error) {
    return s.UserRepository.FindAll()
}
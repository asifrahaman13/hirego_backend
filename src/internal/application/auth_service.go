package application

import "github.com/asifrahaman13/hirego/src/internal/domain"

type AuthService struct {
	AuthRepository domain.AuthRepository
	
}

func NewAuthService(authRepository domain.AuthRepository) *AuthService {
	return &AuthService{AuthRepository: authRepository}
}

func (s *AuthService) Signup(user *domain.User) (interface{}, error) {
	return s.AuthRepository.Signup(user)
}
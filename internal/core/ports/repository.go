package ports

import "github.com/asifrahaman13/hirego/internal/core/domain"

type BaseRepository[T any] interface {
	GetData() ([]*domain.User, error)
	SignUp(*domain.User) (string, error)
	Login(*domain.User) (*domain.AccessToken, error)
	ProtectedRoute(string) (string, error)
	UserInformation(*domain.UserInformation) (string, error)
	GetUserInformation(string) (*domain.UserInformation, error)
}

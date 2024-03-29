package ports

import "github.com/asifrahaman13/hirego/internal/core/domain"

type UserService interface {
	GetAllUsers() ([]*domain.User, error)
	Signup(*domain.User) (string, error)             // Returns a success or a failure message.
	Login(*domain.User) (*domain.AccessToken, error) // Returns a success or a failure message.
	ProtectedRoute(string) (string, error)
	UserInformation(*domain.UserInformation) (string, error)
	GetUserInformation(string) (*domain.UserInformation, error)
	SetUserWrorkInformation(string, *domain.WorkInformation) (string, error)
	GetUserWorkInformation(*domain.UserName)(*domain.WorkInformation, error)
}

type UserRepository interface {
	BaseRepository[*domain.User]
}

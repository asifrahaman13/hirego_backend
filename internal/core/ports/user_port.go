package ports

import "github.com/asifrahaman13/hirego/internal/core/domain"

type UserService interface {
	GetAllUsers() ([]*domain.User, error)
}

type UserRepository interface {
	BaseRepository[*domain.User]
}

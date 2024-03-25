package ports

import "github.com/asifrahaman13/hirego/internal/domain"

type UserService interface {
	GetAllUsers() (map[string]interface{}, error)
}

type UserRepository interface {
	BaseRepository(*domain.User)
}

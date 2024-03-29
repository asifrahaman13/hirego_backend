package ports

import "github.com/asifrahaman13/hirego/internal/core/domain"

type UserService interface {
	Signup(user domain.User) (string, error)
}

type UserRepository interface {
	BaseRepository[domain.User]
}

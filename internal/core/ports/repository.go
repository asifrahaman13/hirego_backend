package ports

import "github.com/asifrahaman13/hirego/internal/core/domain"

// import "github.com/asifrahaman13/hirego/internal/core/domain"

type BaseRepository[T any] interface {
	Create(model T) (string, error)
	GetData() (domain.User, error)
}

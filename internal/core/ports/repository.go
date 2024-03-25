package ports

import "github.com/asifrahaman13/hirego/internal/core/domain"


type BaseRepository[T any] interface{
	// sample

	GetData() ([]*domain.User, error)
}
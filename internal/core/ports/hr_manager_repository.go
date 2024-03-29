package ports

import "github.com/asifrahaman13/hirego/internal/core/domain"

type HRService interface {
	// Signup(*domain.User) (string, error) 
}

type HRRepository interface {
	BaseRepository[domain.HrManager]
}

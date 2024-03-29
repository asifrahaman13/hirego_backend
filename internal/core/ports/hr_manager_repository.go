package ports

import "github.com/asifrahaman13/hirego/internal/core/domain"

type HRService interface {
}

type HRRepository interface {
	BaseRepository[domain.HrManager]
}

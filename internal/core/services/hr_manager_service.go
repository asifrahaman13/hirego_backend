package service

import (
	// "github.com/asifrahaman13/hirego/internal/core/domain"
	"github.com/asifrahaman13/hirego/internal/core/ports"
)

type hrService struct {
	repo ports.HRRepository
}

func InitializeHRService(r ports.HRRepository) *hrService {
	return &hrService{
		repo: r,
	}
}
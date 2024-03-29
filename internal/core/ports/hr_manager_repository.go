package ports

import "github.com/asifrahaman13/hirego/internal/core/domain"

type HRService interface {
	Signup(domain.HrManager) (string, error)
	Login(domain.HrManager) (domain.AccessToken, error)
	SetHrProfileInformation(string, domain.HrProfileInformation) (string, error)
	GetProfileInformation(string) (domain.HrProfileInformation, error)
}

type HRRepository interface {
	BaseRepository[domain.HrManager]
}

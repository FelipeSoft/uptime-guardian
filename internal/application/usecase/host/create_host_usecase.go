package usecase

import (
	"errors"
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
)

type CreateHostUseCase struct {
	repo domain.HostRepository
}

type CreateHostDTO struct {
	IPAddress string `json:"ipAddress"`
	Interval  int64  `json:"interval"`
	Timeout   int64  `json:"timeout"`
}

func NewCreateHostUseCase(repo domain.HostRepository) *CreateHostUseCase {
	return &CreateHostUseCase{
		repo: repo,
	}
}

func (uc *CreateHostUseCase) Execute(dto CreateHostDTO) error {
	if dto.IPAddress == "" {
		return errors.New("at least IP address or URL should be provided")
	}
	u := &domain.Host{
		IPAddress:       dto.IPAddress,
		Interval:  dto.Interval,
		Timeout:   dto.Timeout,
	}
	err := uc.repo.Create(u)
	if err != nil {
		return err
	}
	return nil
}

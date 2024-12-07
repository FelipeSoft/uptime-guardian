package usecase

import (
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
)

type CreateEndpointUseCase struct {
	repo domain.EndpointRepository
}

type CreateEndpointDTO struct {
	URL      string `json:"url"`
	Method   string `json:"method"`
	Interval int64  `json:"interval"`
	Timeout  int64  `json:"timeout"`
}

func NewCreateEndpointUseCase(repo domain.EndpointRepository) *CreateEndpointUseCase {
	return &CreateEndpointUseCase{
		repo: repo,
	}
}

func (uc *CreateEndpointUseCase) Execute(dto CreateEndpointDTO) error {
	u := &domain.Endpoint{
		URL:      dto.URL,
		Method:   dto.Method,
		Interval: dto.Interval,
		Timeout:  dto.Timeout,
	}
	err := uc.repo.Create(u)
	if err != nil {
		return err
	}
	return nil
}

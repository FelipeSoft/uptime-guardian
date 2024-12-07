package usecase

import (
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
)

type CreateEndpointUseCase struct {
	repo domain.EndpointRepository
}

type CreateEndpointDTO struct {
	URL      string `json:"url"       validate:"required"`
	Method   string `json:"method"    validate:"required"`
	Interval int64  `json:"interval"  validate:"required"`
	Timeout  int64  `json:"timeout"   validate:"required"`
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

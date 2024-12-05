package endpoint_usecase

import (
	"errors"
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
)

type CreateEndpointUseCase struct {
	repo domain.EndpointRepository
}

type CreateEndpointDTO struct {
	IPAddress string `json:"ipAddress"`
	URL       string `json:"url"`
	Method    string `json:"method"`
	Interval  int64  `json:"interval"  validate:"required"`
	Timeout   int64  `json:"timeout"   validate:"required"`
}

func NewCreateEndpointUseCase(repo domain.EndpointRepository) *CreateEndpointUseCase {
	return &CreateEndpointUseCase{
		repo: repo,
	}
}

func (uc *CreateEndpointUseCase) Execute(dto CreateEndpointDTO) error {
	if dto.IPAddress == "" && dto.URL == "" {
		return errors.New("at least IP address or URL should be provided")
	}
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

package usecase

import (
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
	"strconv"
)

type UpdateEndpointUseCase struct {
	repo domain.EndpointRepository
}

type UpdateEndpointDTO struct {
	URL      string `json:"url"       validate:"required"`
	Method   string `json:"method"    validate:"required"`
	Interval int64  `json:"interval"  validate:"required"`
	Timeout  int64  `json:"timeout"   validate:"required"`
}

func NewUpdateEndpointUseCase(repo domain.EndpointRepository) *UpdateEndpointUseCase {
	return &UpdateEndpointUseCase{
		repo: repo,
	}
}

func (uc *UpdateEndpointUseCase) Execute(id string, dto UpdateEndpointDTO) error {
	parsedId, err := strconv.Atoi(id)

	if err != nil {
		return err
	}

	u := &domain.Endpoint{
		ID:       uint64(parsedId),
		URL:      dto.URL,
		Method:   dto.Method,
		Interval: dto.Interval,
		Timeout:  dto.Timeout,
	}

	err = uc.repo.Update(u)
	if err != nil {
		return err
	}
	return nil
}

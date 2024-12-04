package usecase

import (
	"strconv"
	"github.com/FelipeSoft/uptime-guardian/internal/http/domain"
)

type UpdateEndpointUseCase struct {
	repo domain.EndpointRepository
}

type UpdateEndpointDTO struct {
	Address  string `json:"address" validate:"required"`
	Interval int64  `json:"interval" validate:"required"`
	Timeout  int64  `json:"timeout" validate:"required"`
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
		ID: uint64(parsedId),
		Address:  dto.Address,
		Interval: dto.Interval,
		Timeout:  dto.Timeout,
	}

	err = uc.repo.Update(u)
	if err != nil {
		return err
	}
	return nil
}

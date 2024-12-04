package usecase

import (
	"strconv"

	"github.com/FelipeSoft/uptime-guardian/internal/http/domain"
)

type DeleteEndpointUseCase struct {
	repo domain.EndpointRepository
}

func NewDeleteEndpointUseCase(repo domain.EndpointRepository) *DeleteEndpointUseCase {
	return &DeleteEndpointUseCase{
		repo: repo,
	}
}

func (uc *DeleteEndpointUseCase) Execute(id string) error {
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = uc.repo.Delete(uint64(parsedId))
	if err != nil {
		return err
	}

	return nil
}
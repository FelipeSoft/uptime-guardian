package usecase

import (
	"strconv"
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
)

type DeleteHostUseCase struct {
	repo domain.HostRepository
}

func NewDeleteHostUseCase(repo domain.HostRepository) *DeleteHostUseCase {
	return &DeleteHostUseCase{
		repo: repo,
	}
}

func (uc *DeleteHostUseCase) Execute(id string) error {
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
package host_usecase

import (
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
	"strconv"
)

type UpdateHostUseCase struct {
	repo domain.HostRepository
}

type UpdateHostDTO struct {
	IPAddress string `json:"ipAddress"`
	Interval  int64  `json:"interval"  validate:"required"`
	Timeout   int64  `json:"timeout"   validate:"required"`
}

func NewUpdateHostUseCase(repo domain.HostRepository) *UpdateHostUseCase {
	return &UpdateHostUseCase{
		repo: repo,
	}
}

func (uc *UpdateHostUseCase) Execute(id string, dto UpdateHostDTO) error {
	parsedId, err := strconv.Atoi(id)

	if err != nil {
		return err
	}

	u := &domain.Host{
		ID:        uint64(parsedId),
		IPAddress: dto.IPAddress,
		Interval:  dto.Interval,
		Timeout:   dto.Timeout,
	}

	err = uc.repo.Update(u)
	if err != nil {
		return err
	}
	return nil
}

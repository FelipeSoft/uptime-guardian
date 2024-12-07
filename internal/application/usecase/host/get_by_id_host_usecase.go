package usecase

import (
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
	"strconv"
)

type GetByIdHostUseCase struct {
	repo domain.HostRepository
}

type GetByIdHostDTO struct {
	ID        uint64 `json:"id"`
	IPAddress string `json:"ipAddress"`
	Interval  int64  `json:"interval"`
	Timeout   int64  `json:"timeout"`
	CreatedAt string `json:"createdAt"`
}

func NewGetByIdHostUseCase(repo domain.HostRepository) *GetByIdHostUseCase {
	return &GetByIdHostUseCase{
		repo: repo,
	}
}

func (uc *GetByIdHostUseCase) Execute(id string) (*GetByIdHostDTO, error) {
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	res, err := uc.repo.GetById(uint64(parsedId))
	if err != nil {
		return nil, err
	}

	output := &GetByIdHostDTO{
		ID:        res.ID,
		IPAddress: res.IPAddress,
		Interval:  res.Interval,
		Timeout:   res.Timeout,
		CreatedAt: res.CreatedAt,
	}

	return output, nil
}

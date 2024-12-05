package host_usecase

import (
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
)

type GetAllHostUseCase struct {
	repo domain.HostRepository
}

type GetAllHostDTO struct {
	ID        uint64 `json:"id"`
	IPAddress string `json:"ipAddress"`
	Interval  int64  `json:"interval"`
	Timeout   int64  `json:"timeout"`
	CreatedAt string `json:"createdAt"`
}

func NewGetAllHostUseCase(repo domain.HostRepository) *GetAllHostUseCase {
	return &GetAllHostUseCase{
		repo: repo,
	}
}

func (uc *GetAllHostUseCase) Execute() ([]*GetAllHostDTO, error) {
	res, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var output []*GetAllHostDTO
	for e := 0; e < len(res); e++ {
		output = append(output, &GetAllHostDTO{
			ID:        res[e].ID,
			IPAddress: res[e].IPAddress,
			Interval:  res[e].Interval,
			Timeout:   res[e].Timeout,
			CreatedAt: res[e].CreatedAt,
		})
	}

	return output, nil
}

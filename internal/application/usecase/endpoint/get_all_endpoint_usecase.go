package endpoint_usecase

import (
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
)

type GetAllEndpointUseCase struct {
	repo domain.EndpointRepository
}

type GetAllEndpointDTO struct {
	ID        uint64 `json:"id"`
	URL       string `json:"url"`
	Method    string `json:"method"`
	Interval  int64  `json:"interval"`
	Timeout   int64  `json:"timeout"`
	CreatedAt string `json:"createdAt"`
}

func NewGetAllEndpointUseCase(repo domain.EndpointRepository) *GetAllEndpointUseCase {
	return &GetAllEndpointUseCase{
		repo: repo,
	}
}

func (uc *GetAllEndpointUseCase) Execute() ([]*GetAllEndpointDTO, error) {
	res, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var output []*GetAllEndpointDTO
	for e := 0; e < len(res); e++ {
		output = append(output, &GetAllEndpointDTO{
			ID:        res[e].ID,
			URL:       res[e].URL,
			Interval:  res[e].Interval,
			Timeout:   res[e].Timeout,
			CreatedAt: res[e].CreatedAt,
		})
	}

	return output, nil
}

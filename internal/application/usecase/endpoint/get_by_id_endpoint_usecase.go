package usecase

import (
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
	"strconv"
)

type GetByIdEndpointUseCase struct {
	repo domain.EndpointRepository
}

type GetByIdEndpointDTO struct {
	ID        uint64 `json:"id"`
	URL       string `json:"url"`
	Method    string `json:"method"`
	Interval  int64  `json:"interval"`
	Timeout   int64  `json:"timeout"`
	CreatedAt string `json:"createdAt"`
}

func NewGetByIdEndpointUseCase(repo domain.EndpointRepository) *GetByIdEndpointUseCase {
	return &GetByIdEndpointUseCase{
		repo: repo,
	}
}

func (uc *GetByIdEndpointUseCase) Execute(id string) (*GetByIdEndpointDTO, error) {
	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	res, err := uc.repo.GetById(uint64(parsedId))
	if err != nil {
		return nil, err
	}

	output := &GetByIdEndpointDTO{
		ID:        res.ID,
		URL:       res.URL,
		Interval:  res.Interval,
		Timeout:   res.Timeout,
		CreatedAt: res.CreatedAt,
	}

	return output, nil
}

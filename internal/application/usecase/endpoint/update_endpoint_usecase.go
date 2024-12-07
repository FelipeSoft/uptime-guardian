package usecase

import (
	"errors"
	"strconv"
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
)

type UpdateEndpointUseCase struct {
	repo domain.EndpointRepository
}

type UpdateEndpointDTO struct {
	URL      *string `json:"url"`     
	Method   *string `json:"method"`  
	Interval *int64  `json:"interval"`
	Timeout  *int64  `json:"timeout"` 
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

	found, err := uc.repo.GetById(uint64(parsedId))
	if err != nil {
		return err
	}
	if found == nil {
		return errors.New("endpoint not found")
	}

	u := &domain.Endpoint{
		ID:       found.ID,
		URL:      chooseString(dto.URL, found.URL),
		Method:   chooseString(dto.Method, found.Method),
		Interval: chooseInt64(dto.Interval, found.Interval),
		Timeout:  chooseInt64(dto.Timeout, found.Timeout),
	}

	err = uc.repo.Update(u)
	if err != nil {
		return err
	}
	return nil
}

func chooseString(newValue *string, existingValue string) string {
	if newValue != nil {
		return *newValue
	}
	return existingValue
}

func chooseInt64(newValue *int64, existingValue int64) int64 {
	if newValue != nil {
		return *newValue
	}
	return existingValue
}

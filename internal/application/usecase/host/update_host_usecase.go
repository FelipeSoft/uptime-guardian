package usecase

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/FelipeSoft/uptime-guardian/internal/domain"
)

type UpdateHostUseCase struct {
	repo domain.HostRepository
}

type UpdateHostDTO struct {
	IPAddress *string `json:"ip_address"`
	Interval  *int64  `json:"interval"`
	Timeout   *int64  `json:"timeout"`
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

	found, err := uc.repo.GetById(uint64(parsedId))
	if err != nil {
		return err
	}

	if found == nil {
		return errors.New("host not found")
	}

	u := &domain.Host{
		ID:        found.ID,
		IPAddress: chooseString(dto.IPAddress, found.IPAddress),
		Interval:  chooseInt64(dto.Interval, found.Interval),
		Timeout:   chooseInt64(dto.Timeout, found.Timeout),
	}
	
	fmt.Println(u)
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

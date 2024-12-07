package usecase

import (
	"fmt"

	"github.com/FelipeSoft/uptime-guardian/internal/domain"
)

type AuthUseCase struct {
	repo domain.UserRepository
}

type LoginUserDTO struct {
	Email    string `json:"email"    validate:"required"`
	Password string `json:"password" validate:"required"`
}

func NewAuthUseCase(repo domain.UserRepository) *AuthUseCase {
	return &AuthUseCase{
		repo: repo,
	}
}

func (uc *AuthUseCase) LoginUser(dto LoginUserDTO) (bool, error) {
	user, err := uc.repo.GetByEmail(dto.Email)
	if err != nil {
		return false, err
	}
	fmt.Println(user)
	return false, nil
}

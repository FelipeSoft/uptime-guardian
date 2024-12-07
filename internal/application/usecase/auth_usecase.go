package usecase

import (
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
)

type AuthUseCase struct {
	repo         domain.UserRepository
	hashable     domain.Hashable
}

type LoginUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthUseCase(repo domain.UserRepository, hashable domain.Hashable) *AuthUseCase {
	return &AuthUseCase{
		repo:         repo,
		hashable:     hashable,
	}
}

func (uc *AuthUseCase) LoginUser(dto LoginUserDTO) (bool, error) {
	user, err := uc.repo.GetByEmail(dto.Email)
	if err != nil {
		return false, err
	}
	authorized := uc.hashable.Compare(dto.Password, user.Password)
	if authorized {
		return true, nil
	}
	return false, nil
}

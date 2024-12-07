package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	auth_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase"
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
)

type AuthHandler struct {
	AuthUseCase *auth_usecase.AuthUseCase
	JwtAdapter  domain.Jwt
}

func NewAuthHandler(AuthUseCase *auth_usecase.AuthUseCase, JwtAdapter domain.Jwt) *AuthHandler {
	return &AuthHandler{
		AuthUseCase: AuthUseCase,
		JwtAdapter:  JwtAdapter,
	}
}

func (ah *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var input auth_usecase.LoginUserDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authorized, err := ah.AuthUseCase.LoginUser(auth_usecase.LoginUserDTO{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if authorized {
		token, err := ah.JwtAdapter.Generate(input.Email)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		cookie := http.Cookie{
			Name:     "UPTIME_GUARDIAN_HTTP",
			Value:    token,
			HttpOnly: true,
			Path:     "/",
			Expires:  time.Now().Add(time.Hour * 1),
			SameSite: http.SameSiteLaxMode,
		}

		if os.Getenv("ENV") == "production" {
			cookie.Secure = true
		}
		if os.Getenv("ENV") != "production" {
			cookie.Secure = false
		}

		http.SetCookie(w, &cookie)
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
}

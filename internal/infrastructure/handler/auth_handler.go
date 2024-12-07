package handler

import (
	"encoding/json"
	"net/http"

	auth_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase"
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
)

type AuthHandler struct {
	AuthUseCase *auth_usecase.AuthUseCase
	JwtAdapter  domain.Jwt
}

func NewAuthHandler(AuthUseCase *auth_usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		AuthUseCase: AuthUseCase,
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
		token := ah.JwtAdapter.Generate(&input.Email)

		cookie := new(http.Cookie)
		cookie.Name = "UPTIME_GUARDIAN_HTTP"
		cookie.Value = token
		cookie.HttpOnly = true

		// possivelmente pode falhar por causa disso
		cookie.Secure = true
		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusAccepted)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
}

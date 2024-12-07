package handler

import (
	"net/http"
	auth_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase"
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
	"github.com/labstack/echo/v4"
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

func (ah *AuthHandler) LoginUser(c echo.Context) error {
	payload, ok := c.Get("payload").(*auth_usecase.LoginUserDTO)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing properties email or password."})
	}
	authorized, err := ah.AuthUseCase.LoginUser(auth_usecase.LoginUserDTO{
		Email:    payload.Email,
		Password: payload.Password,
	})
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	if authorized {
		token := ah.JwtAdapter.Generate(&payload.Email)

		cookie := new(http.Cookie)
		cookie.Name = "UPTIME_GUARDIAN_HTTP"
		cookie.Value = token
		cookie.HttpOnly = true

		// possivelmente pode falhar por causa disso
		cookie.Secure = true

		c.SetCookie(cookie)
		return c.JSON(http.StatusOK, map[string]string{"message": "Logged successfully!"})
	}

	return c.JSON(http.StatusOK, map[string]string{"error": "Invalid email or password"})
}

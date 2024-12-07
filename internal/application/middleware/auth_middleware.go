package middleware

import (
	"net/http"
	"github.com/FelipeSoft/uptime-guardian/internal/domain"
)

type AuthMiddleware struct {
	tokenManager domain.Jwt
}

func NewAuthMiddleware(tokenManager domain.Jwt) *AuthMiddleware {
	return &AuthMiddleware{
		tokenManager: tokenManager,
	}
}

func (m *AuthMiddleware) RequireAuthentication(next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("UPTIME_GUARDIAN_HTTP")
		if r.URL.Path == "/auth/login" {
			next(w, r)
			return
		}
		r.Cookies()
		if err != nil || cookie == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next(w, r)
	})
}

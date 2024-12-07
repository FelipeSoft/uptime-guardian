package middleware

import (
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/labstack/echo/v4"
)

var ExcludeRoutes = []string{
	"/auth/login",
}

func VerifyUserAuthentication(c echo.Context) error {
	if slices.Contains(ExcludeRoutes, c.Path()) {
		return nil
	}

	authCookie, err := c.Cookie("UPTIME_GUARDIAN_HTTP")

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}

	if authCookie.Expires.Before(time.Now()) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Session expired"})
	}

	fmt.Println(authCookie)
	return nil
}

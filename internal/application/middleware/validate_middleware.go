package middleware

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	endpoint_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/endpoint"
	host_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/host"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()

var PayloadRegistry = map[string]map[string]interface{}{
	"/endpoint":     {echo.POST: &endpoint_usecase.CreateEndpointDTO{}},
	"/endpoint/:id": {echo.PUT: &endpoint_usecase.UpdateEndpointDTO{}},
	"/host":     {echo.POST: &host_usecase.CreateHostDTO{}},
	"/host/:id": {echo.PUT: &host_usecase.UpdateHostDTO{}},
}

func matchRoute(path string) string {
	for registeredPath := range PayloadRegistry {
		if registeredPath == path || strings.Contains(registeredPath, ":") && strings.HasPrefix(path, strings.Split(registeredPath, ":")[0]) {
			return registeredPath
		}
	}
	return ""
}

func ValidateRequestBodyDynamic(c echo.Context) error {
	path := c.Path()
	method := c.Request().Method

	route := matchRoute(path)
	if route == "" {
		return nil
	}

	if payloadType, exists := PayloadRegistry[route][method]; exists {
		payload := reflect.New(reflect.TypeOf(payloadType).Elem()).Interface()

		if err := c.Bind(payload); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON format"})
		}

		if err := validate.Struct(payload); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error":   "Validation failed",
				"details": fmt.Sprintf("%v", err),
			})
		}

		c.Set("payload", payload)
	}

	return nil
}

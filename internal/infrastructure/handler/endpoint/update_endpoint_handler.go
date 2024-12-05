package endpoint_handler

import (
	"net/http"
	endpoint_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/endpoint"
	"github.com/labstack/echo/v4"
)

type UpdateEndpointHandler struct {
	UpdateEndpointUseCase *endpoint_usecase.UpdateEndpointUseCase
}

func NewUpdateEndpointHandler(UpdateEndpointUseCase *endpoint_usecase.UpdateEndpointUseCase) *UpdateEndpointHandler {
	return &UpdateEndpointHandler{
		UpdateEndpointUseCase: UpdateEndpointUseCase,
	}
}

func (uc *UpdateEndpointHandler) Execute(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "The 'id' request path value is required"})
	}
	payload, ok := c.Get("payload").(*endpoint_usecase.UpdateEndpointDTO)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Payload not found"})
	}
	if err := uc.UpdateEndpointUseCase.Execute(id, *payload); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update endpoint"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Endpoint updated successfully!"})
}
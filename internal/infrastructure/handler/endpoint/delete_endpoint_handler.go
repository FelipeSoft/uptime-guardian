package handler

import (
	"net/http"
	endpoint_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/endpoint"
	"github.com/labstack/echo/v4"
)

type DeleteEndpointHandler struct {
	DeleteEndpointUseCase *endpoint_usecase.DeleteEndpointUseCase
}

func NewDeleteEndpointHandler(DeleteEndpointUseCase *endpoint_usecase.DeleteEndpointUseCase) *DeleteEndpointHandler {
	return &DeleteEndpointHandler{
		DeleteEndpointUseCase: DeleteEndpointUseCase,
	}
}

func (uc *DeleteEndpointHandler) Execute(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "The 'id' request path value is required"})
	}
	if err := uc.DeleteEndpointUseCase.Execute(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete endpoint"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Endpoint deleted successfully!"})
}
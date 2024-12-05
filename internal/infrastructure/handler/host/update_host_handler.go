package host_handler

import (
	host_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/host"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UpdateHostHandler struct {
	UpdateHostUseCase *host_usecase.UpdateHostUseCase
}

func NewUpdateHostHandler(UpdateHostUseCase *host_usecase.UpdateHostUseCase) *UpdateHostHandler {
	return &UpdateHostHandler{
		UpdateHostUseCase: UpdateHostUseCase,
	}
}

func (uc *UpdateHostHandler) Execute(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "The 'id' request path value is required"})
	}
	payload, ok := c.Get("payload").(*host_usecase.UpdateHostDTO)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Payload not found"})
	}
	if err := uc.UpdateHostUseCase.Execute(id, *payload); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update Host"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Host updated successfully!"})
}

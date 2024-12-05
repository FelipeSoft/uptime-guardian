package host_handler

import (
	"net/http"
	host_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/host"
	"github.com/labstack/echo/v4"
)

type DeleteHostHandler struct {
	DeleteHostUseCase *host_usecase.DeleteHostUseCase
}

func NewDeleteHostHandler(DeleteHostUseCase *host_usecase.DeleteHostUseCase) *DeleteHostHandler {
	return &DeleteHostHandler{
		DeleteHostUseCase: DeleteHostUseCase,
	}
}

func (uc *DeleteHostHandler) Execute(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "The 'id' request path value is required"})
	}
	if err := uc.DeleteHostUseCase.Execute(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete Host"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Host deleted successfully!"})
}
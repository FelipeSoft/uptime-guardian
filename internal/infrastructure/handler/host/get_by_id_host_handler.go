package host_handler

import (
	"net/http"
	host_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/host"
	"github.com/labstack/echo/v4"
)

type GetByIdHostHandler struct {
	GetByIdHostUsecase *host_usecase.GetByIdHostUseCase
}

func NewGetByIdHostHandler(GetByIdHostUsecase *host_usecase.GetByIdHostUseCase) *GetByIdHostHandler {
	return &GetByIdHostHandler{
		GetByIdHostUsecase: GetByIdHostUsecase,
	}
}

func (uc *GetByIdHostHandler) Execute(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "The 'id' request path value is required"})
	}
	res, err := uc.GetByIdHostUsecase.Execute(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed on getting Hosts"})
	}
	return c.JSON(http.StatusOK, res)
}
package handler

import (
	"net/http"
	endpoint_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/endpoint"
	"github.com/labstack/echo/v4"
)

type GetByIdEndpointHandler struct {
	GetByIdEndpointUsecase *endpoint_usecase.GetByIdEndpointUseCase
}

func NewGetByIdEndpointHandler(GetByIdEndpointUsecase *endpoint_usecase.GetByIdEndpointUseCase) *GetByIdEndpointHandler {
	return &GetByIdEndpointHandler{
		GetByIdEndpointUsecase: GetByIdEndpointUsecase,
	}
}

func (uc *GetByIdEndpointHandler) Execute(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "The 'id' request path value is required"})
	}
	res, err := uc.GetByIdEndpointUsecase.Execute(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed on getting endpoints"})
	}
	return c.JSON(http.StatusOK, res)
}
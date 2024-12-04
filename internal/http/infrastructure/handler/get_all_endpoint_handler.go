package handler

import (
	"fmt"
	"net/http"

	"github.com/FelipeSoft/uptime-guardian/internal/http/application/usecase"
	"github.com/labstack/echo/v4"
)

type GetAllEndpointHandler struct {
	GetAllEndpointUsecase *usecase.GetAllEndpointUseCase
}

func NewGetAllEndpointHandler(GetAllEndpointUsecase *usecase.GetAllEndpointUseCase) *GetAllEndpointHandler {
	return &GetAllEndpointHandler{
		GetAllEndpointUsecase: GetAllEndpointUsecase,
	}
}

func (uc *GetAllEndpointHandler) Execute(c echo.Context) error {
	res, err := uc.GetAllEndpointUsecase.Execute()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed on getting endpoints"})
	}
	return c.JSON(http.StatusOK, res)
}
package host_handler

import (
	"fmt"
	"net/http"
	host_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/host"
	"github.com/labstack/echo/v4"
)

type GetAllHostHandler struct {
	GetAllHostUsecase *host_usecase.GetAllHostUseCase
}

func NewGetAllHostHandler(GetAllHostUsecase *host_usecase.GetAllHostUseCase) *GetAllHostHandler {
	return &GetAllHostHandler{
		GetAllHostUsecase: GetAllHostUsecase,
	}
}

func (uc *GetAllHostHandler) Execute(c echo.Context) error {
	res, err := uc.GetAllHostUsecase.Execute()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed on getting Hosts"})
	}
	return c.JSON(http.StatusOK, res)
}
package handler

import (
	"net/http"
	host_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/host"
	"github.com/labstack/echo/v4"
)

type CreateHostHandler struct {
	CreateHostUseCase *host_usecase.CreateHostUseCase
}

func NewCreateHostHandler(CreateHostUseCase *host_usecase.CreateHostUseCase) *CreateHostHandler {
	return &CreateHostHandler{
		CreateHostUseCase: CreateHostUseCase,
	}
}

func (uc *CreateHostHandler) Execute(c echo.Context) error {
	payload, ok := c.Get("payload").(*host_usecase.CreateHostDTO)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Payload not found"})
	}

	if err := uc.CreateHostUseCase.Execute(*payload); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create Host"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Host created successfully!"})
}

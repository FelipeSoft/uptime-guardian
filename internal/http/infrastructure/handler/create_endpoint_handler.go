package handler

import (
	"net/http"

	"github.com/FelipeSoft/uptime-guardian/internal/http/application/usecase"
	"github.com/labstack/echo/v4"
)

type CreateEndpointHandler struct {
	CreateEndpointUseCase *usecase.CreateEndpointUseCase
}

func NewCreateEndpointHandler(CreateEndpointUseCase *usecase.CreateEndpointUseCase) *CreateEndpointHandler {
	return &CreateEndpointHandler{
		CreateEndpointUseCase: CreateEndpointUseCase,
	}
}

func (uc *CreateEndpointHandler) Execute(c echo.Context) error {
	payload, ok := c.Get("payload").(*usecase.CreateEndpointDTO)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Payload not found"})
	}

	if err := uc.CreateEndpointUseCase.Execute(*payload); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create endpoint"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Endpoint created successfully!"})
}

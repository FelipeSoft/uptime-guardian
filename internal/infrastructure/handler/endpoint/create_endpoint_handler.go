package handler

import (
	"encoding/json"
	"net/http"

	endpoint_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/endpoint"
)

type CreateEndpointHandler struct {
	CreateEndpointUseCase *endpoint_usecase.CreateEndpointUseCase
}

func NewCreateEndpointHandler(CreateEndpointUseCase *endpoint_usecase.CreateEndpointUseCase) *CreateEndpointHandler {
	return &CreateEndpointHandler{
		CreateEndpointUseCase: CreateEndpointUseCase,
	}
}

func (uc *CreateEndpointHandler) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var input endpoint_usecase.CreateEndpointDTO
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = uc.CreateEndpointUseCase.Execute(input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

package handler

import (
	"encoding/json"
	"net/http"

	endpoint_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/endpoint"
)

type GetByIdEndpointHandler struct {
	GetByIdEndpointUsecase *endpoint_usecase.GetByIdEndpointUseCase
}

func NewGetByIdEndpointHandler(GetByIdEndpointUsecase *endpoint_usecase.GetByIdEndpointUseCase) *GetByIdEndpointHandler {
	return &GetByIdEndpointHandler{
		GetByIdEndpointUsecase: GetByIdEndpointUsecase,
	}
}

func (uc *GetByIdEndpointHandler) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := uc.GetByIdEndpointUsecase.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

package handler

import (
	"encoding/json"
	"net/http"
	endpoint_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/endpoint"
)

type GetAllEndpointHandler struct {
	GetAllEndpointUsecase *endpoint_usecase.GetAllEndpointUseCase
}

func NewGetAllEndpointHandler(GetAllEndpointUsecase *endpoint_usecase.GetAllEndpointUseCase) *GetAllEndpointHandler {
	return &GetAllEndpointHandler{
		GetAllEndpointUsecase: GetAllEndpointUsecase,
	}
}

func (uc *GetAllEndpointHandler) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	output, err := uc.GetAllEndpointUsecase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

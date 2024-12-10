package handler

import (
	"encoding/json"
	"net/http"
	host_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/host"
)

type GetAllHostHandler struct {
	GetAllHostUsecase *host_usecase.GetAllHostUseCase
}

func NewGetAllHostHandler(GetAllHostUsecase *host_usecase.GetAllHostUseCase) *GetAllHostHandler {
	return &GetAllHostHandler{
		GetAllHostUsecase: GetAllHostUsecase,
	}
}

func (uc *GetAllHostHandler) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	output, err := uc.GetAllHostUsecase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

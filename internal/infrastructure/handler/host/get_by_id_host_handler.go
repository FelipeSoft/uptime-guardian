package handler

import (
	"encoding/json"
	"net/http"

	host_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/host"
)

type GetByIdHostHandler struct {
	GetByIdHostUsecase *host_usecase.GetByIdHostUseCase
}

func NewGetByIdHostHandler(GetByIdHostUsecase *host_usecase.GetByIdHostUseCase) *GetByIdHostHandler {
	return &GetByIdHostHandler{
		GetByIdHostUsecase: GetByIdHostUsecase,
	}
}

func (uc *GetByIdHostHandler) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id := r.PathValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := uc.GetByIdHostUsecase.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

package handler

import (
	"encoding/json"
	"net/http"

	host_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/host"
)

type UpdateHostHandler struct {
	UpdateHostUseCase *host_usecase.UpdateHostUseCase
}

func NewUpdateHostHandler(UpdateHostUseCase *host_usecase.UpdateHostUseCase) *UpdateHostHandler {
	return &UpdateHostHandler{
		UpdateHostUseCase: UpdateHostUseCase,
	}
}

func (uc *UpdateHostHandler) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id := r.PathValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var input host_usecase.UpdateHostDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = uc.UpdateHostUseCase.Execute(id, input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

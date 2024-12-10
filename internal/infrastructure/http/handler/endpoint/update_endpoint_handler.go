package handler

import (
	"encoding/json"
	"net/http"

	endpoint_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/endpoint"
)

type UpdateEndpointHandler struct {
	UpdateEndpointUseCase *endpoint_usecase.UpdateEndpointUseCase
}

func NewUpdateEndpointHandler(UpdateEndpointUseCase *endpoint_usecase.UpdateEndpointUseCase) *UpdateEndpointHandler {
	return &UpdateEndpointHandler{
		UpdateEndpointUseCase: UpdateEndpointUseCase,
	}
}

func (uc *UpdateEndpointHandler) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var input endpoint_usecase.UpdateEndpointDTO
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = uc.UpdateEndpointUseCase.Execute(id, input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

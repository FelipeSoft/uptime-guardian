package handler

import (
	"net/http"

	endpoint_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/endpoint"
)

type DeleteEndpointHandler struct {
	DeleteEndpointUseCase *endpoint_usecase.DeleteEndpointUseCase
}

func NewDeleteEndpointHandler(DeleteEndpointUseCase *endpoint_usecase.DeleteEndpointUseCase) *DeleteEndpointHandler {
	return &DeleteEndpointHandler{
		DeleteEndpointUseCase: DeleteEndpointUseCase,
	}
}

func (uc *DeleteEndpointHandler) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := uc.DeleteEndpointUseCase.Execute(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

package handler

import (
	"net/http"

	host_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/host"
)

type DeleteHostHandler struct {
	DeleteHostUseCase *host_usecase.DeleteHostUseCase
}

func NewDeleteHostHandler(DeleteHostUseCase *host_usecase.DeleteHostUseCase) *DeleteHostHandler {
	return &DeleteHostHandler{
		DeleteHostUseCase: DeleteHostUseCase,
	}
}

func (uc *DeleteHostHandler) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return 
	}
	id := r.PathValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := uc.DeleteHostUseCase.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

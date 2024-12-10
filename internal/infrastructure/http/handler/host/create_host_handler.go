package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	host_usecase "github.com/FelipeSoft/uptime-guardian/internal/application/usecase/host"
)

type CreateHostHandler struct {
	CreateHostUseCase *host_usecase.CreateHostUseCase
}

func NewCreateHostHandler(CreateHostUseCase *host_usecase.CreateHostUseCase) *CreateHostHandler {
	return &CreateHostHandler{
		CreateHostUseCase: CreateHostUseCase,
	}
}

func (uc *CreateHostHandler) Execute(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var input host_usecase.CreateHostDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = uc.CreateHostUseCase.Execute(input)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

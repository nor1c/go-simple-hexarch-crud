package http

import (
	"encoding/json"
	"net/http"
	"sha/pkg/usecase"
)

type UserHandler struct {
	UserUseCase usecase.UserUseCase
}

func (h *UserHandler) HandleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserUseCase.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(users)
}

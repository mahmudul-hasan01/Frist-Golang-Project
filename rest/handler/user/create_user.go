package user

import (
	"back-end/repo"
	"back-end/util"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var newUser repo.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	createdUser, err := h.userRepo.Create(newUser)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, createdUser, http.StatusCreated)
}

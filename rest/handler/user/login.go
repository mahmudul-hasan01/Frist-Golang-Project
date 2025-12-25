package user

import (
	"back-end/util"
	"encoding/json"
	"net/http"
	"strconv"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	var loginRequest LoginRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginRequest)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	usr, err := h.userRepo.Find(loginRequest.Email, loginRequest.Password)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if usr == nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	accessToken, err := util.CreateJwt(h.cnf.JwtSecretKey, util.Payload{
		Sub:         strconv.Itoa(usr.ID),
		FirstName:   usr.FirstName,
		LastName:    usr.LastName,
		Email:       usr.Email,
		IsShopOwner: usr.IsShopOwner,
		Role:        usr.Role,
	})
	if err != nil {
		http.Error(w, "Could not create accessToken", http.StatusInternalServerError)
		return
	}

	util.SendData(w, map[string]interface{}{
		"user":        usr,
		"accessToken": accessToken,
	}, http.StatusOK)
}

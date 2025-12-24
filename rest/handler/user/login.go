package user

import (
	"back-end/config"
	"back-end/database"
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

	usr := database.FindUserByEmail(loginRequest.Email, loginRequest.Password)
	if usr == nil {
		http.Error(w, "Invalid email or password", http.StatusBadRequest)
		return
	}

	cnf := config.GetConfig()

	accessToken, err := util.CreateJwt(cnf.JwtSecretKey, util.Payload{
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
	}, http.StatusCreated)
}

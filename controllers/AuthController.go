package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Jose-Guerrero-Developer/twittorbackend/authentication"

	"github.com/Jose-Guerrero-Developer/twittorbackend/utils"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
)

// AuthLogin login
func AuthLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var User models.User
	var Auth models.Auth

	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "001", "Error getting data", err.Error())
		return
	}

	status := Auth.Authentication(&User)

	if !status {
		w.WriteHeader(http.StatusUnauthorized)
		response := utils.ResponseErrorJWT{
			Code:        "006",
			Message:     "Authentication",
			Description: "Access credentials are inconsistent",
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	token, err := authentication.GenerateToken(&User)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := utils.ResponseErrorJWT{
			Code:        "007",
			Message:     "Authentication",
			Description: "Impossible to generate access token. " + err.Error(),
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	w.WriteHeader(http.StatusCreated)
	signature := authentication.JWT{
		Token: token,
	}
	json.NewEncoder(w).Encode(&signature)

	expiresToken := time.Now().Add(1 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expiresToken,
	})
}

package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Jose-Guerrero-Developer/twittorbackend/authentication"

	"github.com/Jose-Guerrero-Developer/twittorbackend/utils"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
)

/*AuthController auth controller */
type AuthController struct{}

/*Sign return user login */
func (Controller *AuthController) Sign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var User models.User
	var Auth models.Auth
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		utils.ResponseFailed(w, "001", "Error getting data", err.Error(), http.StatusBadRequest)
		return
	}
	status := Auth.Sign(&User)
	if !status {
		utils.ResponseFailed(w, "006", "Authentication", "Access credentials are inconsistent", http.StatusUnauthorized)
		return
	}
	token, err := authentication.GenerateToken(&User)
	if err != nil {
		utils.ResponseFailed(w, "007", "Authentication", "Impossible to generate access token. "+err.Error(), http.StatusBadRequest)
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

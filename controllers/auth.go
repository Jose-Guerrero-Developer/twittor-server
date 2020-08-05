package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
)

/*Auth auth controller */
type Auth struct {
	galex.Controller
}

/*Sign return user login */
func (Controller *Auth) Sign(w http.ResponseWriter, r *http.Request) {
	var User models.User
	var Auth models.Auth
	var JWT models.JWT
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		Controller.Response().Failed("001", "Error getting data", err.Error(), http.StatusBadRequest)
		return
	}
	status := Auth.Sign(&User)
	if !status {
		Controller.Response().Failed("006", "Authentication", "Access credentials are inconsistent", http.StatusUnauthorized)
		return
	}
	token, err := JWT.GenerateToken(&User)
	if err != nil {
		Controller.Response().Failed("007", "Authentication", "Impossible to generate access token. "+err.Error(), http.StatusBadRequest)
		return
	}
	Controller.Response().Success(bson.M{"token": token}, http.StatusCreated)
}

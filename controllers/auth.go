package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/devJGuerrero/twittor-server/galex/response"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/devJGuerrero/twittor-server/models"
)

/*Auth auth controller */
type Auth struct{}

/*Sign return user login */
func (Controller *Auth) Sign(w http.ResponseWriter, r *http.Request) {
	var User models.User
	var Auth models.Auth
	var JWT models.JWT
	var GalexResponse response.Driver
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		GalexResponse.Failed("001", "Error getting data", err.Error(), http.StatusBadRequest)
		return
	}
	status := Auth.Sign(&User)
	if !status {
		GalexResponse.Failed("006", "Authentication", "Access credentials are inconsistent", http.StatusUnauthorized)
		return
	}
	token, err := JWT.GenerateToken(&User)
	if err != nil {
		GalexResponse.Failed("007", "Authentication", "Impossible to generate access token. "+err.Error(), http.StatusBadRequest)
		return
	}
	GalexResponse.Success(bson.M{"token": token}, http.StatusCreated)
}

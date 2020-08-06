package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/response"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
)

/*User user controller */
type User struct{}

/*Store return stores a user in a database */
func (Controller *User) Store(w http.ResponseWriter, r *http.Request) {
	var GalexResponse response.Driver
	var User models.User
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		GalexResponse.Failed("001", "Error getting data", err.Error(), http.StatusBadRequest)
		return
	}
	if len(User.Email) == 0 {
		GalexResponse.Failed("002", "Required field", "Email is a required field", http.StatusBadRequest)
		return
	}
	if len(User.Password) <= 5 {
		GalexResponse.Failed("003", "Field length", "Password is a field that must be at least 6 characters", http.StatusBadRequest)
		return
	}
	if User.ExistsEmail() {
		GalexResponse.Failed("004", "Duplicate data", "Email account is registered", http.StatusConflict)
		return
	}
	status, id, message := User.Store()
	if !status {
		GalexResponse.Failed("005", "Database transaction", "Error storing user data. "+message.Error(), http.StatusConflict)
		return
	}
	GalexResponse.Success(bson.M{
		"id": id,
	}, http.StatusCreated)
}

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/utils"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
)

/*UserController user controller */
type UserController struct{}

/*Store return stores a user in a database */
func (Controller *UserController) Store(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var User models.User
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		utils.ResponseFailed(w, "001", "Error getting data", err.Error(), http.StatusBadRequest)
		return
	}
	if len(User.Email) == 0 {
		utils.ResponseFailed(w, "002", "Required field", "Email is a required field", http.StatusBadRequest)
		return
	}
	if len(User.Password) <= 5 {
		utils.ResponseFailed(w, "003", "Field length", "Password is a field that must be at least 6 characters", http.StatusBadRequest)
		return
	}
	if User.Exists() {
		utils.ResponseFailed(w, "004", "Duplicate data", "Email account is registered", http.StatusConflict)
		return
	}
	status, id, message := User.Insert()
	if !status {
		utils.ResponseFailed(w, "005", "Database transaction", "Error storing user data. "+message.Error(), http.StatusConflict)
		return
	}
	utils.Response(w, bson.M{
		"_id": id,
	}, http.StatusCreated)
}

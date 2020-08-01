package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/utils"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
)

/*UserController user controller */
type UserController struct{}

/*Store return stores a user in a database */
func (Controller *UserController) Store(w http.ResponseWriter, r *http.Request) {
	var User models.User

	err := json.NewDecoder(r.Body).Decode(&User)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "001", "Error getting data", err.Error())
		return
	}

	if len(User.Email) == 0 {
		utils.ResponseError(w, http.StatusBadRequest, "002", "Required field", "Email is a required field")
		return
	}

	if len(User.Password) <= 5 {
		utils.ResponseError(w, http.StatusBadRequest, "003", "Field length", "Password is a field that must be at least 6 characters")
		return
	}

	if User.Exists() {
		utils.ResponseError(w, http.StatusConflict, "004", "Duplicate data", "Email account is registered")
		return
	}

	status, id, message := User.Insert()
	if !status {
		utils.ResponseError(w, http.StatusConflict, "005", "Database transaction", "Error storing user data. "+message.Error())
		return
	}

	data := map[string]string{"id": id}
	utils.Response(w, http.StatusCreated, data)
}

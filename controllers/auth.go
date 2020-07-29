package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/utils"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
)

// AuthLogin login
func AuthLogin(w http.ResponseWriter, r *http.Request) {
	var User models.User
	var Auth models.Auth

	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "001", "Error getting data", err.Error())
		return
	}

	status := Auth.Authentication(&User)
	data := map[string]string{}
	statusHTTP := http.StatusOK

	if !status {
		statusHTTP = http.StatusUnauthorized
		data["code"] = "006"
		data["message"] = "Authentication"
		data["description"] = "Access credentials are inconsistent"
	}

	if status {
		data["accessToken"] = "token de acceso"
		data["name"] = User.Name
		data["email"] = User.Email
		data["lastName"] = User.LastName
		data["dateBirth"] = User.DateBirth.String()
		data["avatar"] = User.Avatar
		data["banner"] = User.Banner
		data["biography"] = User.Biography
		data["location"] = User.Location
		data["website"] = User.Website
	}

	utils.Response(w, statusHTTP, data)
}

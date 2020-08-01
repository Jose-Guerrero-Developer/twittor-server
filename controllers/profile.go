package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/utils"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
)

/*ProfileController profile controller */
type ProfileController struct{}

/*Get return a user profile by id */
func (Controller *ProfileController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var Profile models.ProfileModel
	ID := r.URL.Query().Get("id")
	err := Profile.Get(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := utils.ResponseErrorJWT{
			Code:        "008",
			Message:     "Error obtaining data",
			Description: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode(Profile)
}

/*Update Update user profile in session */
func (Controller *ProfileController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var Profile models.ProfileModel

	err := json.NewDecoder(r.Body).Decode(&Profile)
	if err != nil {
		response := utils.ResponseErrorJWT{
			Code:        "001",
			Message:     "Error getting data",
			Description: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	status := true
	var ProfileUpdate bson.M
	status, ProfileUpdate, err = Profile.Update()
	if err != nil || !status {
		response := utils.ResponseErrorJWT{
			Code:        "009",
			Message:     "Error updating resource",
			Description: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}
	json.NewEncoder(w).Encode(ProfileUpdate)
}

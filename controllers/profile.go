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
	var Profile models.Profile
	ID := r.URL.Query().Get("id")
	if len(ID) <= 0 {
		utils.ResponseFailed(w, "011", "Required parameter", "It is necessary to send in the application an id profile", http.StatusBadRequest)
		return
	}
	err := Profile.Get(ID)
	if err != nil {
		utils.ResponseFailed(w, "008", "Error obtaining data", err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(Profile)
}

/*Update Update user profile in session */
func (Controller *ProfileController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Profile models.Profile
	err := json.NewDecoder(r.Body).Decode(&Profile)
	if err != nil {
		utils.ResponseFailed(w, "001", "Error getting data", err.Error(), http.StatusBadRequest)
		return
	}
	status := true
	var ProfileUpdate bson.M
	status, ProfileUpdate, err = Profile.Update()
	if err != nil || !status {
		utils.ResponseFailed(w, "009", "Error updating resource", err.Error(), http.StatusBadRequest)
		return
	}
	utils.Response(w, ProfileUpdate, http.StatusOK)
}

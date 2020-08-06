package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/response"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
)

/*Profile profile controller */
type Profile struct{}

/*Get return a user profile by id */
func (Controller *Profile) Get(w http.ResponseWriter, r *http.Request) {
	var GalexResponse response.Driver
	var Profile models.Profile
	ID := r.URL.Query().Get("id")
	if len(ID) <= 0 {
		GalexResponse.Failed("011", "Required parameter", "It is necessary to send in the application an id profile", http.StatusBadRequest)
		return
	}
	err := Profile.Get(ID)
	if err != nil {
		GalexResponse.Failed("008", "Error obtaining data", err.Error(), http.StatusBadRequest)
		return
	}
	GalexResponse.Success(Profile, http.StatusOK)
}

/*Update Update user profile in session */
func (Controller *Profile) Update(w http.ResponseWriter, r *http.Request) {
	var GalexResponse response.Driver
	var Profile models.Profile
	err := json.NewDecoder(r.Body).Decode(&Profile)
	if err != nil {
		GalexResponse.Failed("001", "Error getting data", err.Error(), http.StatusBadRequest)
		return
	}
	status := true
	var ProfileUpdate bson.M
	status, ProfileUpdate, err = Profile.Update()
	if err != nil || !status {
		GalexResponse.Failed("009", "Error updating resource", err.Error(), http.StatusBadRequest)
		return
	}
	GalexResponse.Success(ProfileUpdate, http.StatusOK)
}

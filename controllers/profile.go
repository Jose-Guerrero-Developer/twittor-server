package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/devJGuerrero/twittor-server/galex/response"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/devJGuerrero/twittor-server/models"
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

/*GetAvatar Get avatar from a profile */
func (Controller *Profile) GetAvatar(w http.ResponseWriter, r *http.Request) {
	var Profile models.Profile
	var GalexResponse response.Driver

	params := mux.Vars(r)
	IDProfile := params["id"]
	if err := Profile.Get(IDProfile); err != nil {
		GalexResponse.Failed("012", "Resource in the found", "Profile id not found", http.StatusNotFound)
		return
	}
	open, err := os.Open("uploads/avatars/" + Profile.Avatar)
	if err != nil {
		GalexResponse.Failed("015", "Error opening writing path", err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/jpeg; image/png")
	_, err = io.Copy(w, open)
	if err != nil {
		GalexResponse.Failed("016", "Error copying image to destination folder", err.Error(), http.StatusBadRequest)
		return
	}
}

/*GetBanner Get banner from a profile */
func (Controller *Profile) GetBanner(w http.ResponseWriter, r *http.Request) {
	var Profile models.Profile
	var GalexResponse response.Driver

	params := mux.Vars(r)
	IDProfile := params["id"]
	if err := Profile.Get(IDProfile); err != nil {
		GalexResponse.Failed("012", "Resource in the found", "Profile id not found", http.StatusNotFound)
		return
	}
	open, err := os.Open("uploads/banners/" + Profile.Banner)
	if err != nil {
		GalexResponse.Failed("015", "Error opening writing path", err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/jpeg; image/png")
	_, err = io.Copy(w, open)
	if err != nil {
		GalexResponse.Failed("016", "Error copying image to destination folder", err.Error(), http.StatusBadRequest)
		return
	}
}

/*Update Update user profile in session */
func (Controller *Profile) Update(w http.ResponseWriter, r *http.Request) {
	var ProfileUpdate bson.M
	var Profile models.Profile
	var GalexResponse response.Driver

	params := mux.Vars(r)
	IDProfile := params["id"]
	err := json.NewDecoder(r.Body).Decode(&Profile)
	if err != nil {
		GalexResponse.Failed("001", "Error getting data", err.Error(), http.StatusBadRequest)
		return
	}
	status := true
	status, ProfileUpdate, err = Profile.Update(IDProfile)
	if err != nil || !status {
		GalexResponse.Failed("009", "Error updating resource", err.Error(), http.StatusBadRequest)
		return
	}
	GalexResponse.Success(ProfileUpdate, http.StatusOK)
}

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
)

/*Tweet tweet controller */
type Tweet struct {
	galex.Controller
}

/*GetProfile return all tweets in a profile */
func (Controller *Tweet) GetProfile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDProfile := params["id"]
	if len(IDProfile) < 1 {
		Controller.Response().Failed("011", "Required parameter", "It is necessary to send in the application an id profile", http.StatusBadRequest)
		return
	}
	var Profile models.Profile
	Profile.ID, _ = primitive.ObjectIDFromHex(IDProfile)
	if status := Profile.ExistsID(); !status {
		Controller.Response().Failed("012", "Resource in the found", "Profile id not found", http.StatusNotFound)
		return
	}
	var Tweet models.Tweet
	if results, status := Tweet.GetProfile(Profile.ID); status && results != nil {
		Controller.Response().Success(results, http.StatusOK)
	} else {
		Controller.Response().Success(bson.M{}, http.StatusOK)
	}
}

/*Store tweetrecord invoegen in database */
func (Controller *Tweet) Store(w http.ResponseWriter, r *http.Request) {
	var Tweet models.Tweet
	var Profile models.Profile
	if err := json.NewDecoder(r.Body).Decode(&Tweet); err != nil {
		Controller.Response().Failed("001", "Error getting data", err.Error(), http.StatusBadRequest)
		return
	}
	if Tweet.IDProfile = Profile.GetID(); !Profile.ExistsID() {
		Controller.Response().Failed("012", "Resource in the found", "Profile id not found", http.StatusNotFound)
		return
	}
	if len(Tweet.Message) <= 0 {
		Controller.Response().Failed("011", "Required parameter", "It is necessary to send in the application an message", http.StatusBadRequest)
		return
	}
	if status, err := Tweet.Store(); !status || err != nil {
		Controller.Response().Failed("005", "Database transaction", "Error storing tweet data. "+err.Error(), http.StatusConflict)
		return
	}
	Controller.Response().Success(bson.M{}, http.StatusCreated)
}

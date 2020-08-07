package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/response"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
)

/*Tweet tweet controller */
type Tweet struct{}

/*Get Returns all tweets */
func (Controller *Tweet) Get(w http.ResponseWriter, r *http.Request) {
	var Tweet models.Tweet
	var GalexResponse response.Driver
	if results, err := Tweet.Get(); err == nil {
		GalexResponse.Success(results, http.StatusOK)
	} else {
		GalexResponse.Success(bson.M{}, http.StatusOK)
	}
}

/*GetID return tweet id */
func (Controller *Tweet) GetID(w http.ResponseWriter, r *http.Request) {
	var Tweet models.Tweet
	var GalexResponse response.Driver

	params := mux.Vars(r)
	IDTweet := params["id"]
	if len(IDTweet) < 1 {
		GalexResponse.Failed("011", "Required parameter", "It is necessary to send in the application an id profile", http.StatusBadRequest)
		return
	}

	result, err := Tweet.GetID(IDTweet)
	if err != nil {
		GalexResponse.Success(bson.M{}, http.StatusNotFound)
		return
	}
	GalexResponse.Success(result, http.StatusOK)
}

/*GetProfile return all tweets in a profile */
func (Controller *Tweet) GetProfile(w http.ResponseWriter, r *http.Request) {
	var GalexResponse response.Driver
	params := mux.Vars(r)
	IDProfile := params["id"]
	if len(IDProfile) < 1 {
		GalexResponse.Failed("011", "Required parameter", "It is necessary to send in the application an id profile", http.StatusBadRequest)
		return
	}
	var Profile models.Profile
	Profile.ID, _ = primitive.ObjectIDFromHex(IDProfile)
	if status := Profile.ExistsID(); !status {
		GalexResponse.Failed("012", "Resource in the found", "Profile id not found", http.StatusNotFound)
		return
	}
	var Tweet models.Tweet
	if results, status := Tweet.GetProfile(Profile.ID); status && results != nil {
		GalexResponse.Success(results, http.StatusOK)
	} else {
		GalexResponse.Success(bson.M{}, http.StatusOK)
	}
}

/*Store tweetrecord invoegen in database */
func (Controller *Tweet) Store(w http.ResponseWriter, r *http.Request) {
	var GalexResponse response.Driver
	var Tweet models.Tweet
	var Profile models.Profile
	if err := json.NewDecoder(r.Body).Decode(&Tweet); err != nil {
		GalexResponse.Failed("001", "Error getting data", err.Error(), http.StatusBadRequest)
		return
	}
	if Tweet.IDProfile = Profile.GetID(); !Profile.ExistsID() {
		GalexResponse.Failed("012", "Resource in the found", "Profile id not found", http.StatusNotFound)
		return
	}
	if len(Tweet.Message) <= 0 {
		GalexResponse.Failed("011", "Required parameter", "It is necessary to send in the application an message", http.StatusBadRequest)
		return
	}
	if status, err := Tweet.Store(); !status || err != nil {
		GalexResponse.Failed("005", "Database transaction", "Error storing tweet data. "+err.Error(), http.StatusConflict)
		return
	}
	GalexResponse.Success(bson.M{}, http.StatusCreated)
}

/*Delete Remove tweet */
func (Controller *Tweet) Delete(w http.ResponseWriter, r *http.Request) {
	var GalexResponse response.Driver
	var Tweet models.Tweet

	params := mux.Vars(r)
	IDTweet := params["id"]
	if err := Tweet.Delete(IDTweet); err == nil {
		GalexResponse.Success(bson.M{}, http.StatusOK)
	} else {
		GalexResponse.Failed("013", "Remove resource", err.Error(), http.StatusBadRequest)
	}
}

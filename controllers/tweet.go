package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/utils/request"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/response"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
)

/*Tweet Controller Tweets */
type Tweet struct{}

/*Get Returns all tweets */
func (Controller *Tweet) Get(w http.ResponseWriter, r *http.Request) {
	var Tweet models.Tweet
	var GalexRequest request.Driver
	var GalexResponse response.Driver

	if data, err := Tweet.Get(); err == nil {
		GalexRequest.AddHeader("X-Total-Count", strconv.Itoa(len(data)))
		GalexResponse.Success(data, http.StatusOK)
		return
	}
	GalexResponse.Success(bson.M{}, http.StatusOK)
}

/*GetID Return tweet ID */
func (Controller *Tweet) GetID(w http.ResponseWriter, r *http.Request) {
	var Tweet models.Tweet
	var GalexResponse response.Driver

	params := mux.Vars(r)
	IDTweet := params["id"]
	if data, err := Tweet.GetID(IDTweet); err == nil {
		GalexResponse.Success(data, http.StatusOK)
		return
	}
	GalexResponse.Success(bson.M{}, http.StatusNotFound)
}

/*GetProfile Return all tweets in a profile */
func (Controller *Tweet) GetProfile(w http.ResponseWriter, r *http.Request) {
	var Tweet models.Tweet
	var Profile models.Profile
	var GalexRequest request.Driver
	var GalexResponse response.Driver

	params := mux.Vars(r)
	IDProfile := params["id"]
	if exists := Profile.ExistsID(IDProfile); !exists {
		GalexResponse.Failed("012", "Resource in the found", "Profile id not found", http.StatusNotFound)
		return
	}
	if data, err := Tweet.GetProfile(IDProfile); data != nil && err == nil {
		GalexRequest.AddHeader("X-Total-Count", strconv.Itoa(len(data)))
		GalexResponse.Success(data, http.StatusOK)
		return
	}
	GalexResponse.Success(bson.M{}, http.StatusOK)
}

/*Store Store a tweet in the database */
func (Controller *Tweet) Store(w http.ResponseWriter, r *http.Request) {
	var Tweet models.Tweet
	var Profile models.Profile
	var GalexResponse response.Driver

	if err := json.NewDecoder(r.Body).Decode(&Tweet); err != nil {
		GalexResponse.Failed("001", "Error getting data", err.Error(), http.StatusBadRequest)
		return
	}
	if len(Tweet.Message) <= 0 {
		GalexResponse.Failed("011", "Required parameter", "It is necessary to send in the application an message", http.StatusBadRequest)
		return
	}
	if exists := Profile.ExistsID(Tweet.IDProfile.Hex()); !exists {
		GalexResponse.Failed("012", "Resource in the found", "Profile id not found", http.StatusNotFound)
		return
	}
	if status, err := Tweet.Store(); !status || err != nil {
		GalexResponse.Failed("005", "Database transaction", "Error storing tweet data. "+err.Error(), http.StatusConflict)
		return
	}
	GalexResponse.Success(bson.M{}, http.StatusCreated)
}

/*Update Update a tweet in the database */
func (Controller *Tweet) Update(w http.ResponseWriter, r *http.Request) {
	var Tweet models.Tweet
	var GalexResponse response.Driver

	params := mux.Vars(r)
	IDTweet := params["id"]
	if err := json.NewDecoder(r.Body).Decode(&Tweet); err != nil {
		GalexResponse.Failed("001", "Error getting data", err.Error(), http.StatusBadRequest)
		return
	}
	if result, err := Tweet.Update(IDTweet); err == nil {
		GalexResponse.Success(result, http.StatusOK)
	} else {
		GalexResponse.Failed("009", "Error updating resource", err.Error(), http.StatusBadRequest)
	}
}

/*Delete Delete a tweet in the database */
func (Controller *Tweet) Delete(w http.ResponseWriter, r *http.Request) {
	var Tweet models.Tweet
	var GalexResponse response.Driver

	params := mux.Vars(r)
	IDTweet := params["id"]
	if err := Tweet.Delete(IDTweet); err == nil {
		GalexResponse.Success(bson.M{}, http.StatusOK)
	} else {
		GalexResponse.Failed("013", "Remove resource", err.Error(), http.StatusBadRequest)
	}
}

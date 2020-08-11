package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/response"
	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/utils/request"
	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

/*Follow Controller Followers */
type Follow struct{}

/*GetProfile Returns all followers profile */
func (Controller *Follow) GetProfile(w http.ResponseWriter, r *http.Request) {
	var Followers models.Follow
	var GalexRequest request.Driver
	var GalexResponse response.Driver

	params := mux.Vars(r)
	IDProfile := params["id"]
	search := r.URL.Query().Get("search")
	if data, err := Followers.GetProfile(IDProfile, search); data != nil && err == nil {
		GalexRequest.AddHeader("X-Total-Count", strconv.Itoa(len(data)))
		GalexResponse.Success(data, http.StatusOK)
		return
	}
	GalexResponse.Success(bson.M{}, http.StatusOK)
}

/*GetFollowed Returns all followed profile */
func (Controller *Follow) GetFollowed(w http.ResponseWriter, r *http.Request) {
	var Followers models.Follow
	var GalexRequest request.Driver
	var GalexResponse response.Driver

	search := r.URL.Query().Get("search")
	if data, err := Followers.GetFollowed(search); data != nil && err == nil {
		GalexRequest.AddHeader("X-Total-Count", strconv.Itoa(len(data)))
		GalexResponse.Success(data, http.StatusOK)
		return
	}
	GalexResponse.Success(bson.M{}, http.StatusOK)
}

/*Exists There is a relationship with the profile to be followed */
func (Controller *Follow) Exists(w http.ResponseWriter, r *http.Request) {
	var Followers models.Follow
	var Profile models.Profile
	var GalexResponse response.Driver

	params := mux.Vars(r)
	IDProfile := params["idProfile"]
	IDFollow := params["idFollow"]
	if exists := Profile.ExistsID(IDProfile); !exists {
		GalexResponse.Failed("012", "Resource in the found", "Profile id not found", http.StatusNotFound)
		return
	}
	if exists := Profile.ExistsID(IDFollow); !exists {
		GalexResponse.Failed("012", "Resource in the found", "Follow id not found", http.StatusNotFound)
		return
	}
	if exists := Followers.Exists(IDProfile, IDFollow); !exists {
		GalexResponse.Success(bson.M{"status": false}, http.StatusOK)
		return
	}
	GalexResponse.Success(bson.M{"status": true}, http.StatusOK)
}

/*Store Store a follow in the database */
func (Controller *Follow) Store(w http.ResponseWriter, r *http.Request) {
	var Followers models.Follow
	var Profile models.Profile
	var GalexResponse response.Driver

	if err := json.NewDecoder(r.Body).Decode(&Followers); err != nil {
		GalexResponse.Failed("001", "Error getting data", err.Error(), http.StatusBadRequest)
		return
	}
	if exists := Profile.ExistsID(Followers.IDProfile.Hex()); !exists {
		GalexResponse.Failed("012", "Resource in the found", "Profile id not found", http.StatusNotFound)
		return
	}
	if exists := Profile.ExistsID(Followers.IDFollow.Hex()); !exists {
		GalexResponse.Failed("012", "Resource in the found", "Follow id not found", http.StatusNotFound)
		return
	}
	if exists := Followers.Exists(Followers.IDProfile.Hex(), Followers.IDFollow.Hex()); exists {
		GalexResponse.Failed("004", "Duplicate data", "There is a relationship with the profile to be followed", http.StatusConflict)
		return
	}
	if Followers.IDFollow.Hex() == models.UID {
		GalexResponse.Failed("018", "Operation denied", "It's not valid to follow yourself", http.StatusConflict)
		return
	}
	if status, err := Followers.Store(); !status || err != nil {
		GalexResponse.Failed("005", "Database transaction", "Error storing tweet data. "+err.Error(), http.StatusConflict)
		return
	}
	GalexResponse.Success(bson.M{}, http.StatusCreated)
}

/*Delete Delete a follow in the database */
func (Controller *Follow) Delete(w http.ResponseWriter, r *http.Request) {
	var Followers models.Follow
	var Profile models.Profile
	var GalexResponse response.Driver

	if err := json.NewDecoder(r.Body).Decode(&Followers); err != nil {
		GalexResponse.Failed("001", "Error getting data", err.Error(), http.StatusBadRequest)
		return
	}
	if exists := Profile.ExistsID(Followers.IDProfile.Hex()); !exists {
		GalexResponse.Failed("012", "Resource in the found", "Profile id not found", http.StatusNotFound)
		return
	}
	if exists := Profile.ExistsID(Followers.IDFollow.Hex()); !exists {
		GalexResponse.Failed("012", "Resource in the found", "Follow id not found", http.StatusNotFound)
		return
	}
	if err := Followers.Delete(Followers.IDProfile.Hex(), Followers.IDFollow.Hex()); err == nil {
		GalexResponse.Success(bson.M{}, http.StatusOK)
	} else {
		GalexResponse.Failed("013", "Remove resource", err.Error(), http.StatusBadRequest)
	}
}

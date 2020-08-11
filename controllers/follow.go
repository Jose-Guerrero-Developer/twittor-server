package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/response"
	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*Follow Controller Followers */
type Follow struct{}

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

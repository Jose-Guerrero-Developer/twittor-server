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
	var Follow models.Follow
	var Profile models.Profile
	var GalexResponse response.Driver

	if err := json.NewDecoder(r.Body).Decode(&Follow); err != nil {
		GalexResponse.Failed("001", "Error getting data", err.Error(), http.StatusBadRequest)
		return
	}
	if exists := Profile.ExistsID(Follow.IDProfile.Hex()); !exists {
		GalexResponse.Failed("012", "Resource in the found", "Profile id not found", http.StatusNotFound)
		return
	}
	if exists := Profile.ExistsID(Follow.IDFollow.Hex()); !exists {
		GalexResponse.Failed("012", "Resource in the found", "Follow id not found", http.StatusNotFound)
		return
	}
	if exists := Follow.Exists(Follow.IDProfile.Hex(), Follow.IDFollow.Hex()); exists {
		GalexResponse.Failed("004", "Duplicate data", "There is a relationship with the profile to be followed", http.StatusConflict)
		return
	}
	if status, err := Follow.Store(); !status || err != nil {
		GalexResponse.Failed("005", "Database transaction", "Error storing tweet data. "+err.Error(), http.StatusConflict)
		return
	}
	GalexResponse.Success(bson.M{}, http.StatusCreated)
}

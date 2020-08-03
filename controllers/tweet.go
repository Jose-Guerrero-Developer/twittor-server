package controllers

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Jose-Guerrero-Developer/twittorbackend/utils"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
)

/*Tweet tweet controller */
type Tweet struct{}

/*Store tweetrecord invoegen in database */
func (Controller *Tweet) Store(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Tweet models.Tweet
	var Profile models.Profile
	if err := json.NewDecoder(r.Body).Decode(&Tweet); err != nil {
		utils.ResponseFailed(w, "001", "Error getting data", err.Error(), http.StatusBadRequest)
		return
	}
	if Tweet.IDProfile = Profile.GetID(); !Profile.ExistsID() {
		utils.ResponseFailed(w, "012", "Resource in the found", "Profile id not found", http.StatusNotFound)
		return
	}
	if len(Tweet.Message) <= 0 {
		utils.ResponseFailed(w, "011", "Required parameter", "It is necessary to send in the application an message", http.StatusBadRequest)
		return
	}
	if status, err := Tweet.Store(); !status || err != nil {
		utils.ResponseFailed(w, "005", "Database transaction", "Error storing tweet data. "+err.Error(), http.StatusConflict)
		return
	}
	utils.Response(w, bson.M{}, http.StatusCreated)
}

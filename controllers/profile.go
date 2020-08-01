package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Jose-Guerrero-Developer/twittorbackend/utils"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
)

/*ProfileController profile controller */
type ProfileController struct{}

/*Get return a user profile by id */
func (profileController *ProfileController) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var Profile models.Profile
	ID := r.URL.Query().Get("id")
	err := Profile.Get(ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := utils.ResponseErrorJWT{
			Code:        "008",
			Message:     "Error obtaining data",
			Description: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	json.NewEncoder(w).Encode(Profile)
}

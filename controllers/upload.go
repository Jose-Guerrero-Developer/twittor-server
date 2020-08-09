package controllers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/response"
)

/*Upload Controller Uploads */
type Upload struct{}

/*Avatar Upload avatar image */
func (Controller *Upload) Avatar(w http.ResponseWriter, r *http.Request) {
	var open *os.File
	var Profile models.Profile
	var GalexResponse response.Driver

	params := mux.Vars(r)
	IDProfile := params["id"]
	if exists := Profile.ExistsID(IDProfile); !exists {
		GalexResponse.Failed("012", "Resource in the found", "Profile id not found", http.StatusNotFound)
		return
	}
	avatar, hanlder, err := r.FormFile("avatar")
	if err != nil {
		GalexResponse.Failed("014", "Error uploading image", err.Error(), http.StatusBadRequest)
		return
	}
	extension := strings.Split(hanlder.Filename, ".")[1]
	path := "uploads/avatars/"
	fileName := IDProfile + "." + extension
	file := path + fileName
	if Profile.Get(IDProfile); Profile.Avatar != "" {
		os.Remove(path + Profile.Avatar)
	}
	open, err = os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		GalexResponse.Failed("015", "Error opening writing path", err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(open, avatar)
	if err != nil {
		GalexResponse.Failed("016", "Error copying image to destination folder", err.Error(), http.StatusBadRequest)
		return
	}
	Profile.Avatar = fileName
	if status, _, err := Profile.Update(IDProfile); !status || err != nil {
		GalexResponse.Failed("009", "Error updating resource", err.Error(), http.StatusBadRequest)
		return
	}
	GalexResponse.Success(bson.M{}, http.StatusCreated)
}

/*Banner Upload banner image */
func (Controller *Upload) Banner(w http.ResponseWriter, r *http.Request) {
	var open *os.File
	var Profile models.Profile
	var GalexResponse response.Driver

	params := mux.Vars(r)
	IDProfile := params["id"]
	if exists := Profile.ExistsID(IDProfile); !exists {
		GalexResponse.Failed("012", "Resource in the found", "Profile id not found", http.StatusNotFound)
		return
	}
	banner, hanlder, err := r.FormFile("banner")
	if err != nil {
		GalexResponse.Failed("014", "Error uploading image", err.Error(), http.StatusBadRequest)
		return
	}
	extension := strings.Split(hanlder.Filename, ".")[1]
	path := "uploads/banners/"
	fileName := IDProfile + "." + extension
	file := path + fileName
	if Profile.Get(IDProfile); Profile.Banner != "" {
		os.Remove(path + Profile.Banner)
	}
	open, err = os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		GalexResponse.Failed("015", "Error opening writing path", err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(open, banner)
	if err != nil {
		GalexResponse.Failed("016", "Error copying image to destination folder", err.Error(), http.StatusBadRequest)
		return
	}
	Profile.Banner = fileName
	if status, _, err := Profile.Update(IDProfile); !status || err != nil {
		GalexResponse.Failed("009", "Error updating resource", err.Error(), http.StatusBadRequest)
		return
	}
	GalexResponse.Success(bson.M{}, http.StatusCreated)
}

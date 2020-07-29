package authentication

import (
	"os"
	"time"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
	"github.com/dgrijalva/jwt-go"
)

// JWT struct
type JWT struct {
	Token string `json:"token,omitempty"`
}

// GenerateToken generate access token
func GenerateToken(_User *models.User) (string, error) {
	secret := []byte(os.Getenv("SECRET"))

	payload := jwt.MapClaims{
		"_id":       _User.ID.Hex(),
		"name":      _User.Name,
		"lastName":  _User.LastName,
		"dateBirth": _User.DateBirth,
		"email":     _User.Email,
		"avatar":    _User.Avatar,
		"banner":    _User.Banner,
		"biography": _User.Biography,
		"location":  _User.Location,
		"webSite":   _User.Website,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	}

	signature := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := signature.SignedString(secret)

	if err != nil {
		return token, err
	}

	return token, nil
}

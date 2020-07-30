package authentication

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
	"github.com/dgrijalva/jwt-go"
)

// Email email user
var Email string

// IDUser id user
var IDUser string

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

// ValidateToken validate token access
func ValidateToken(token string) (*models.ClaimJWT, bool, string, error) {
	claims := &models.ClaimJWT{}
	secret := []byte(os.Getenv("SECRET"))

	splitToken := strings.Split(token, "Bearer ")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Invalid token format")
	}

	token = strings.TrimSpace(splitToken[1])
	signature, err := jwt.ParseWithClaims(token, claims, func(_token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if !signature.Valid {
		return claims, false, string(""), errors.New("Invalid token format")
	}

	if err == nil {
		_User := new(models.User)
		_User.Email = claims.Email
		exists := _User.CheckExists()
		if exists {
			IDUser = claims.ID.Hex()
			Email = claims.Email
		}
		return claims, exists, IDUser, nil
	}

	return claims, false, string(""), err
}

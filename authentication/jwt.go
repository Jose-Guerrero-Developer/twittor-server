package authentication

import (
	"errors"
	"strings"
	"time"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
	"github.com/dgrijalva/jwt-go"
)

/*JWT structure for managing authentication (JWT) */
type JWT struct {
	Token string `json:"token,omitempty"`
}

/*GenerateToken return JWT generated access token */
func GenerateToken(User *models.User) (string, error) {
	var Galex galex.Driver
	secret := []byte(Galex.Configs().Get("APP_SECRET"))
	payload := jwt.MapClaims{
		"_id":       User.ID.Hex(),
		"name":      User.Name,
		"lastName":  User.LastName,
		"dateBirth": User.DateBirth,
		"email":     User.Email,
		"avatar":    User.Avatar,
		"banner":    User.Banner,
		"biography": User.Biography,
		"location":  User.Location,
		"webSite":   User.Website,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	}
	signature := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := signature.SignedString(secret)
	if err != nil {
		return token, err
	}
	return token, nil
}

/*ValidateToken return access token validation */
func ValidateToken(token string) (*models.ClaimJWT, bool, string, error) {
	var Galex galex.Driver
	Claims := &models.ClaimJWT{}
	secret := []byte(Galex.Configs().Get("APP_SECRET"))
	splitToken := strings.Split(token, "Bearer ")
	if len(splitToken) != 2 {
		return Claims, false, string(""), errors.New("Invalid token format")
	}
	token = strings.TrimSpace(splitToken[1])
	signature, err := jwt.ParseWithClaims(token, Claims, func(_token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if !signature.Valid {
		return Claims, false, string(""), errors.New("Invalid token format")
	}
	if err == nil {
		var User models.User
		User.Email = Claims.Email
		exists := User.ExistsEmail()
		if exists {
			models.IDUser = Claims.ID.Hex()
			models.Email = Claims.Email
		}
		return Claims, exists, models.IDUser, nil
	}
	return Claims, false, string(""), err
}

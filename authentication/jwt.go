package authentication

import (
	"errors"
	"strings"
	"time"

	"github.com/Jose-Guerrero-Developer/twittorbackend/configs"

	"github.com/Jose-Guerrero-Developer/twittorbackend/models"
	"github.com/dgrijalva/jwt-go"
)

/*JWT structure for managing authentication (JWT) */
type JWT struct {
	Token string `json:"token,omitempty"`
}

/*GenerateToken return JWT generated access token */
func GenerateToken(_User *models.User) (string, error) {
	var Configs configs.Driver
	secret := []byte(Configs.Get("APP_SECRET"))
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

/*ValidateToken return access token validation */
func ValidateToken(token string) (*models.ClaimJWT, bool, string, error) {
	claims := &models.ClaimJWT{}
	var Configs configs.Driver
	secret := []byte(Configs.Get("APP_SECRET"))

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
		var UserModel models.User
		UserModel.Email = claims.Email
		exists := UserModel.Exists()
		if exists {
			models.IDUser = claims.ID.Hex()
			models.Email = claims.Email
		}
		return claims, exists, models.IDUser, nil
	}
	return claims, false, string(""), err
}

package models

import (
	"errors"
	"strings"
	"time"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/configuration"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*JWT claims structure for the JWT model */
type JWT struct {
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Email string             `json:"email"`
	jwt.StandardClaims
}

/*GenerateToken return JWT generated access token */
func (Model *JWT) GenerateToken(User *User) (string, error) {
	var GalexConfigs configuration.Driver
	secret := []byte(GalexConfigs.Get("APP_SECRET"))
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
func (Model *JWT) ValidateToken(token string) (*JWT, bool, string, error) {
	var GalexConfigs configuration.Driver
	Claims := &JWT{}
	secret := []byte(GalexConfigs.Get("APP_SECRET"))
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
		var User User
		User.Email = Claims.Email
		exists := User.ExistsEmail()
		if exists {
			IDProfile = Claims.ID.Hex()
			EmailProfile = Claims.Email
		}
		return Claims, exists, IDProfile, nil
	}
	return Claims, false, string(""), err
}

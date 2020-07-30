package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ClaimJWT claims jwt
type ClaimJWT struct {
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Email string             `json:"email"`
	jwt.StandardClaims
}

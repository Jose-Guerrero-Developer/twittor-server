package models

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/Jose-Guerrero-Developer/twittorbackend/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct model
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	LastName  string             `bson:"lastName" json:"lastName"`
	DateBirth time.Time          `bson:"dateBirth" json:"dateBirth"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar"`
	Banner    string             `bson:"banner" json:"banner"`
	Biography string             `bson:"biography" json:"biography"`
	Location  string             `bson:"location" json:"location"`
	Website   string             `bson:"website" json:"website"`
}

// CheckExists validate if exists user
func (_User *User) CheckExists() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := database.Connection.Database("twittor")
	users := db.Collection("users")

	err := users.FindOne(ctx, bson.M{"email": _User.Email}).Decode(&_User)

	if err != nil {
		return false
	}

	return true
}

// Insert recod in database
func (_User *User) Insert() (bool, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := database.Connection.Database("twittor")
	users := db.Collection("users")

	_User.Password, _ = encryptPassword(_User.Password)

	record, err := users.InsertOne(ctx, _User)

	if err != nil {
		return false, "", err
	}

	id, _ := record.InsertedID.(primitive.ObjectID)
	return true, id.String(), err
}

func encryptPassword(password string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

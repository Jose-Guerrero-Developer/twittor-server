package models

import (
	"context"
	"time"

	"github.com/Jose-Guerrero-Developer/twittorbackend/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*User structure to manage user model */
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

/*Exists returns if a user exists */
func (UserModel *User) Exists() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	Users := ORM.Collection("users")
	err := Users.FindOne(ctx, bson.M{"email": UserModel.Email}).Decode(&UserModel)

	if err != nil {
		return false
	}
	return true
}

/*Insert stores a user in a database */
func (UserModel *User) Insert() (bool, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	Users := ORM.Collection("users")
	UserModel.Password, _ = utils.EncryptPassword(UserModel.Password)

	record, err := Users.InsertOne(ctx, UserModel)
	if err != nil {
		return false, "", err
	}

	id, _ := record.InsertedID.(primitive.ObjectID)
	return true, id.String(), err
}

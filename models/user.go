package models

import (
	"context"
	"time"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*User structure to manage user model */
type User struct {
	galex.Model
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
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

/*ExistsID returns if a user exists */
func (Model *User) ExistsID() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	Users := ORM.Collection("users")
	err := Users.FindOne(ctx, bson.M{"_id": Model.ID}).Decode(&Model)
	if err != nil {
		return false
	}
	return true
}

/*ExistsEmail returns if a user exists */
func (Model *User) ExistsEmail() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	Users := ORM.Collection("users")
	err := Users.FindOne(ctx, bson.M{"email": Model.Email}).Decode(&Model)
	if err != nil {
		return false
	}
	return true
}

/*Insert stores a user in a database */
func (Model *User) Insert() (bool, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	Users := ORM.Collection("users")
	Model.Password, _ = Model.Utils().Bcrypt.EncryptPassword(Model.Password)
	record, err := Users.InsertOne(ctx, Model)
	if err != nil {
		return false, "", err
	}
	id, _ := record.InsertedID.(primitive.ObjectID)
	return true, id.String(), err
}

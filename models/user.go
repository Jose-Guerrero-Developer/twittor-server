package models

import (
	"context"
	"time"

	"github.com/devJGuerrero/twittor-server/galex/utils/bcrypt"

	"github.com/devJGuerrero/twittor-server/galex/database/helpers"

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

/*ExistsID returns if a user exists */
func (Model *User) ExistsID(IDProfile string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var Users = helpers.EstablishDriver("users")
	ID, _ := primitive.ObjectIDFromHex(IDProfile)
	err := Users.FindOne(ctx, bson.M{"_id": bson.M{"$eq": ID}}).Decode(&Model)
	if err != nil {
		return false
	}
	return true
}

/*ExistsEmail returns if a user exists */
func (Model *User) ExistsEmail() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var Users = helpers.EstablishDriver("users")
	err := Users.FindOne(ctx, bson.M{"email": Model.Email}).Decode(&Model)
	if err != nil {
		return false
	}
	return true
}

/*Store stores a user in a database */
func (Model *User) Store() (bool, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var GalexBcrypt bcrypt.Driver
	var Users = helpers.EstablishDriver("users")
	Model.Password, _ = GalexBcrypt.EncryptPassword(Model.Password)
	record, err := Users.InsertOne(ctx, Model)
	if err != nil {
		return false, "", err
	}
	id, _ := record.InsertedID.(primitive.ObjectID)
	return true, id.Hex(), err
}

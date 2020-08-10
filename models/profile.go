package models

import (
	"context"
	"time"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/database/helpers"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Profile management structure for profile template */
type Profile struct {
	User
}

/*Get return profile data */
func (Model *Profile) Get(ID string) error {
	objID, _ := primitive.ObjectIDFromHex(ID)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var Users = helpers.EstablishDriver("users")
	condition := bson.M{
		"_id": objID,
	}
	err := Users.FindOne(ctx, condition).Decode(&Model.User)
	Model.Password = ""
	if err != nil {
		return err
	}
	return nil
}

/*Update Update user profile in session */
func (Model *Profile) Update(ID string) (bool, bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var Users = helpers.EstablishDriver("users")
	IDProfile, _ := primitive.ObjectIDFromHex(ID)
	Record := make(map[string]interface{})
	if Model.Name != "" {
		Record["name"] = Model.Name
	}
	if Model.LastName != "" {
		Record["lastName"] = Model.LastName
	}
	if !Model.DateBirth.IsZero() {
		Record["dateBirth"] = Model.DateBirth
	}
	if Model.Avatar != "" {
		Record["avatar"] = Model.Avatar
	}
	if Model.Banner != "" {
		Record["banner"] = Model.Banner
	}
	if Model.Biography != "" {
		Record["biography"] = Model.Biography
	}
	if Model.Location != "" {
		Record["location"] = Model.Location
	}
	if Model.Website != "" {
		Record["website"] = Model.Website
	}
	ActionUpdate := bson.M{
		"$set": Record,
	}
	condition := bson.M{
		"_id": bson.M{"$eq": IDProfile},
	}
	var ModelUpdate bson.M
	_, err := Users.UpdateOne(ctx, condition, ActionUpdate)
	if err != nil {
		return false, ModelUpdate, err
	}
	err = Users.FindOne(ctx, bson.M{"_id": IDProfile}).Decode(&ModelUpdate)
	delete(ModelUpdate, "password")
	if err != nil {
		return false, ModelUpdate, err
	}
	return true, ModelUpdate, nil
}

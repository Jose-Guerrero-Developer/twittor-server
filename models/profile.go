package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Profile management structure for profile template */
type Profile struct {
	User
}

/*Get return profile data */
func (ProfileModel *Profile) Get(ID string) error {
	objID, _ := primitive.ObjectIDFromHex(ID)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	Users := ORM.Collection("users")
	condition := bson.M{
		"_id": objID,
	}

	err := Users.FindOne(ctx, condition).Decode(&ProfileModel.User)
	ProfileModel.Password = ""

	if err != nil {
		return err
	}
	return nil
}

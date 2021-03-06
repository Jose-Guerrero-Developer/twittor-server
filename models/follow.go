package models

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/devJGuerrero/twittor-server/galex/database/helpers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Follow Model Followers */
type Follow struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	IDProfile primitive.ObjectID `bson:"_id_profile" json:"idProfile"`
	IDFollow  primitive.ObjectID `bson:"_id_follow" json:"idFollow"`
}

/*GetProfile Returns all followers profile */
func (Model *Follow) GetProfile(idProfile string, search string) ([]*User, error) {
	var data []*User
	var Profiles = helpers.EstablishDriver("users")
	var Followers = helpers.EstablishDriver("followers")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	IDProfile, _ := primitive.ObjectIDFromHex(idProfile)
	cursor, err := Followers.Find(ctx, bson.M{"_id_profile": bson.M{"$eq": IDProfile}})
	if err != nil {
		return data, err
	}
	for cursor.Next(ctx) {
		var follow *Follow
		var user *User
		if err := cursor.Decode(&follow); err != nil {
			log.Println("Impossible transforms data follow")
			continue
		}
		if err := Profiles.FindOne(ctx, bson.M{"_id": bson.M{"$eq": follow.IDFollow}, "name": bson.M{"$regex": `(?i)` + search}}).Decode(&user); err != nil {
			continue
		}
		user.Password = ""
		data = append(data, user)
	}
	return data, nil
}

/*GetFollowed Returns all followed profile */
func (Model *Follow) GetFollowed(search string) ([]*User, error) {
	var data []*User
	var Profiles = helpers.EstablishDriver("users")
	var Followers = helpers.EstablishDriver("followers")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	IDProfile, _ := primitive.ObjectIDFromHex(UID)
	cursor, err := Followers.Find(ctx, bson.M{"_id_follow": bson.M{"$eq": IDProfile}})
	if err != nil {
		return data, err
	}
	for cursor.Next(ctx) {
		var follow *Follow
		var user *User
		if err := cursor.Decode(&follow); err != nil {
			log.Println("Impossible transforms data follow")
			continue
		}
		if err := Profiles.FindOne(ctx, bson.M{"_id": bson.M{"$eq": follow.IDProfile}, "name": bson.M{"$regex": `(?i)` + search}}).Decode(&user); err != nil {
			continue
		}
		user.Password = ""
		data = append(data, user)
	}
	return data, nil
}

/*Exists There is a relationship between followers */
func (Model *Follow) Exists(idProfile string, idFollow string) bool {
	var Followers = helpers.EstablishDriver("followers")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	IDProfile, _ := primitive.ObjectIDFromHex(idProfile)
	IDFollow, _ := primitive.ObjectIDFromHex(idFollow)
	if err := Followers.FindOne(ctx, bson.M{"_id_profile": bson.M{"$eq": IDProfile}, "_id_follow": bson.M{"$eq": IDFollow}}); err.Err() != nil {
		return false
	}
	return true
}

/*Store Store a follow in the database */
func (Model *Follow) Store() (bool, error) {
	var Followers = helpers.EstablishDriver("followers")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if _, err := Followers.InsertOne(ctx, Model); err != nil {
		return false, err
	}
	return true, nil
}

/*Delete Delete a follow in the database */
func (Model *Follow) Delete(idProfile string, idFollow string) error {
	var Followers = helpers.EstablishDriver("followers")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	IDProfile, _ := primitive.ObjectIDFromHex(idProfile)
	IDFollow, _ := primitive.ObjectIDFromHex(idFollow)
	if result, err := Followers.DeleteOne(ctx, bson.M{"_id_profile": bson.M{"$eq": IDProfile}, "_id_follow": bson.M{"$eq": IDFollow}}); err != nil || result.DeletedCount < 1 {
		return errors.New("It is not possible to remove the resource")
	}
	return nil
}

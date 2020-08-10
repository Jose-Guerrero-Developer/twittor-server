package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/database/helpers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Follow Model Followers */
type Follow struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	IDProfile primitive.ObjectID `bson:"_id_profile" json:"idProfile"`
	IDFollow  primitive.ObjectID `bson:"_id_follow" json:"idFollow"`
}

/*Exists There is a relationship between followers */
func (Model *Follow) Exists(idProfile string, idFollow string) bool {
	var Follow = helpers.EstablishDriver("follow")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	IDProfile, _ := primitive.ObjectIDFromHex(idProfile)
	IDFollow, _ := primitive.ObjectIDFromHex(idFollow)
	if err := Follow.FindOne(ctx, bson.M{"_id_profile": bson.M{"$eq": IDProfile}, "_id_follow": bson.M{"$eq": IDFollow}}); err.Err() != nil {
		return false
	}
	return true
}

/*Store Store a follow in the database */
func (Model *Follow) Store() (bool, error) {
	var Follow = helpers.EstablishDriver("follow")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if _, err := Follow.InsertOne(ctx, Model); err != nil {
		return false, err
	}
	return true, nil
}

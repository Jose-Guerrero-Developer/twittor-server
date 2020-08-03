package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Tweet structure to manage tweet model */
type Tweet struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	IDProfile primitive.ObjectID `bson:"_id_profile" json:"idProfile"`
	Message   string             `bson:"message" json:"message"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
}

/*Store register a tweet in the database */
func (Model *Tweet) Store() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	Tweets := ORM.Collection("tweets")
	Model.CreatedAt = time.Now()
	_, err := Tweets.InsertOne(ctx, Model)
	if err != nil {
		return false, err
	}
	return true, nil
}

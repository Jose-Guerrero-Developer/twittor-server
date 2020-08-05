package models

import (
	"context"
	"log"
	"time"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Tweet structure to manage tweet model */
type Tweet struct {
	galex.Model
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	IDProfile primitive.ObjectID `bson:"_id_profile" json:"idProfile"`
	Message   string             `bson:"message" json:"message"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
}

/*GetProfile return all tweets in a profile */
func (Model *Tweet) GetProfile(IDProfile primitive.ObjectID) ([]*Tweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var results []*Tweet
	var record Tweet
	Tweets := ORM.Collection("tweets")
	filter := bson.M{
		"_id_profile": IDProfile,
	}
	opts := options.Find()
	opts.SetLimit(Model.Utils().Pagination.GetCount())
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
	opts.SetSkip((Model.Utils().Pagination.GetPage() - 1) * Model.Utils().Pagination.GetCount())
	cursor, err := Tweets.Find(ctx, filter, opts)
	if err != nil {
		return results, false
	}
	for cursor.Next(ctx) {
		if err := cursor.Decode(&record); err != nil {
			log.Println("Impossible transforms data tweet GetProfile")
			continue
		}
		results = append(results, &record)
	}
	return results, true
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

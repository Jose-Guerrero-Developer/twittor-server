package models

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/utils/pagination"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/database/helpers"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Tweet structure to manage tweet model */
type Tweet struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	IDProfile primitive.ObjectID `bson:"_id_profile" json:"idProfile"`
	Message   string             `bson:"message" json:"message"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
}

/*Get Return all tweet */
func (Model *Tweet) Get() ([]*Tweet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var GalexORM helpers.Driver
	var GalexPagination pagination.Driver
	var results []*Tweet
	Tweets := GalexORM.Collection("tweets")
	opts := options.Find()
	opts.SetLimit(GalexPagination.GetCount())
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
	opts.SetSkip((GalexPagination.GetPage() - 1) * GalexPagination.GetCount())
	cursor, err := Tweets.Find(ctx, bson.M{}, opts)
	if err != nil {
		return results, err
	}
	for cursor.Next(ctx) {
		var record Tweet
		if err := cursor.Decode(&record); err != nil {
			log.Println("Impossible transforms data tweet GetProfile")
			continue
		}
		results = append(results, &record)
	}
	GalexPagination.AddHeader("X-Total-Count", strconv.Itoa(len(results)))
	return results, nil
}

/*GetID Return tweet id */
func (Model *Tweet) GetID(IDTweet string) (*Tweet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var GalexORM helpers.Driver
	Tweets := GalexORM.Collection("tweets")
	id, _ := primitive.ObjectIDFromHex(IDTweet)
	filter := bson.M{
		"_id": id,
	}
	err := Tweets.FindOne(ctx, filter).Decode(&Model)
	if err != nil {
		return Model, err
	}
	return Model, nil
}

/*GetProfile return all tweets in a profile */
func (Model *Tweet) GetProfile(IDProfile primitive.ObjectID) ([]*Tweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var GalexORM helpers.Driver
	var GalexPagination pagination.Driver
	var results []*Tweet
	Tweets := GalexORM.Collection("tweets")
	filter := bson.M{
		"_id_profile": IDProfile,
	}
	opts := options.Find()
	opts.SetLimit(GalexPagination.GetCount())
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
	opts.SetSkip((GalexPagination.GetPage() - 1) * GalexPagination.GetCount())
	cursor, err := Tweets.Find(ctx, filter, opts)
	if err != nil {
		return results, false
	}
	for cursor.Next(ctx) {
		var record Tweet
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

	var GalexORM helpers.Driver
	Tweets := GalexORM.Collection("tweets")
	Model.CreatedAt = time.Now()
	_, err := Tweets.InsertOne(ctx, Model)
	if err != nil {
		return false, err
	}
	return true, nil
}

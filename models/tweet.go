package models

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/utils/pagination"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/database/helpers"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Tweet Model Tweets */
type Tweet struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	IDProfile primitive.ObjectID `bson:"_id_profile" json:"idProfile"`
	Message   string             `bson:"message" json:"message"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
}

/*Get Returns all tweets */
func (Model *Tweet) Get() ([]*Tweet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var data []*Tweet
	var GalexORM helpers.Driver
	var GalexPagination pagination.Driver

	Tweets := GalexORM.Collection("tweets")
	opts := options.Find()
	opts.SetLimit(GalexPagination.GetCount())
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
	opts.SetSkip((GalexPagination.GetPage() - 1) * GalexPagination.GetCount())
	cursor, err := Tweets.Find(ctx, bson.M{}, opts)
	if err != nil {
		return data, err
	}
	for cursor.Next(ctx) {
		var record Tweet
		if err := cursor.Decode(&record); err != nil {
			log.Println("Impossible transforms data tweet get profile")
			continue
		}
		data = append(data, &record)
	}
	GalexPagination.AddHeader("X-Total-Count", strconv.Itoa(len(data)))
	return data, nil
}

/*GetID Return tweet ID */
func (Model *Tweet) GetID(IDTweet string) (*Tweet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var GalexORM helpers.Driver
	Tweets := GalexORM.Collection("tweets")
	ID, _ := primitive.ObjectIDFromHex(IDTweet)
	filter := bson.M{"_id": bson.M{"$eq": ID}}
	if err := Tweets.FindOne(ctx, filter).Decode(&Model); err != nil {
		return Model, err
	}
	return Model, nil
}

/*GetProfile Return all tweets in a profile */
func (Model *Tweet) GetProfile(IDProfile string) ([]*Tweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var data []*Tweet
	var GalexORM helpers.Driver
	var GalexPagination pagination.Driver

	Tweets := GalexORM.Collection("tweets")
	ID, _ := primitive.ObjectIDFromHex(IDProfile)
	filter := bson.M{"_id_profile": bson.M{"$eq": ID}}
	opts := options.Find()
	opts.SetLimit(GalexPagination.GetCount())
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
	opts.SetSkip((GalexPagination.GetPage() - 1) * GalexPagination.GetCount())
	cursor, err := Tweets.Find(ctx, filter, opts)
	if err != nil {
		return data, false
	}
	for cursor.Next(ctx) {
		var record Tweet
		if err := cursor.Decode(&record); err != nil {
			log.Println("Impossible transforms data tweet GetProfile")
			continue
		}
		data = append(data, &record)
	}
	return data, true
}

/*Store Store a tweet in the database */
func (Model *Tweet) Store() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var GalexORM helpers.Driver

	Tweets := GalexORM.Collection("tweets")
	Model.CreatedAt = time.Now()
	if _, err := Tweets.InsertOne(ctx, Model); err != nil {
		return false, err
	}
	return true, nil
}

/*Update Update a tweet in the database */
func (Model *Tweet) Update(IDTweet string) (*Tweet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var GalexORM helpers.Driver

	ID, _ := primitive.ObjectIDFromHex(IDTweet)
	Tweets := GalexORM.Collection("tweets")
	data := make(map[string]interface{})
	if !Model.IDProfile.IsZero() {
		data["_id_profile"] = Model.IDProfile
	}
	if Model.Message != "" {
		data["message"] = Model.Message
	}
	filter := bson.M{"_id": bson.M{"$eq": ID}}
	if _, err := Tweets.UpdateOne(ctx, filter, bson.M{"$set": data}); err != nil {
		return Model, err
	}
	if err := Tweets.FindOne(ctx, filter).Decode(&Model); err != nil {
		return Model, err
	}
	return Model, nil
}

/*Delete Delete a tweet in the database */
func (Model *Tweet) Delete(IDTweet string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var GalexORM helpers.Driver

	Tweets := GalexORM.Collection("tweets")
	ID, _ := primitive.ObjectIDFromHex(IDTweet)
	filter := bson.M{"_id": bson.M{"$eq": ID}}
	if result, err := Tweets.DeleteOne(ctx, filter); err != nil || result.DeletedCount < 1 {
		return errors.New("It is not possible to remove the resource")
	}
	return nil
}

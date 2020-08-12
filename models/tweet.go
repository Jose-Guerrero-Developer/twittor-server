package models

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/devJGuerrero/twittor-server/galex/database/helpers"

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

/*TweetFollow Model TweetsFollow */
type TweetFollow struct {
	User struct {
		ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
		Name     string             `bson:"name" json:"name"`
		LastName string             `bson:"lastName" json:"lastName"`
		Email    string             `bson:"email" json:"email"`
	} `json:"user"`
	Tweet struct {
		ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
		Message   string             `bson:"message" json:"message"`
		CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
	} `json:"tweet"`
}

/*Get Returns all tweets */
func (Model *Tweet) Get() ([]*Tweet, error) {
	var data []*Tweet
	var Tweets = helpers.EstablishDriver("tweets")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	cursor, err := Tweets.Find(ctx, bson.M{})
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
	return data, nil
}

/*GetID Return tweet ID */
func (Model *Tweet) GetID(IDTweet string) (*Tweet, error) {
	var Tweets = helpers.EstablishDriver("tweets")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(IDTweet)
	if err := Tweets.FindOne(ctx, bson.M{"_id": bson.M{"$eq": ID}}).Decode(&Model); err != nil {
		return Model, err
	}
	return Model, nil
}

/*GetProfile Return all tweets in a profile */
func (Model *Tweet) GetProfile(IDProfile string) ([]*Tweet, error) {
	var data []*Tweet
	var Tweets = helpers.EstablishDriver("tweets")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(IDProfile)
	cursor, err := Tweets.Find(ctx, bson.M{"_id_profile": bson.M{"$eq": ID}})
	if err != nil {
		return data, err
	}
	for cursor.Next(ctx) {
		var record Tweet
		if err := cursor.Decode(&record); err != nil {
			log.Println("Impossible transforms data tweet GetProfile")
			continue
		}
		data = append(data, &record)
	}
	return data, nil
}

/*GetFollow Return all tweets follow */
func (Model *Tweet) GetFollow() ([]*TweetFollow, error) {
	var data []*TweetFollow
	var Followers = helpers.EstablishDriver("followers")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	skip := (1 - 1) * 20
	IDProfile, _ := primitive.ObjectIDFromHex(UID)
	pepiline := make([]bson.M, 0)
	pepiline = append(pepiline, bson.M{"$match": bson.M{"_id_profile": IDProfile}})
	pepiline = append(pepiline, bson.M{
		"$lookup": bson.M{
			"from":         "users",
			"as":           "user",
			"foreignField": "_id",
			"localField":   "_id_follow",
		},
	})
	pepiline = append(pepiline, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",
			"as":           "tweet",
			"foreignField": "_id_profile",
			"localField":   "_id_follow",
		},
	})
	pepiline = append(pepiline, bson.M{"$unwind": "$user"})
	pepiline = append(pepiline, bson.M{"$unwind": "$tweet"})
	pepiline = append(pepiline, bson.M{"$sort": bson.M{"tweets.created_at": -1}})
	pepiline = append(pepiline, bson.M{"$skip": skip})
	pepiline = append(pepiline, bson.M{"$limit": 20})
	cursor, _ := Followers.Aggregate(ctx, pepiline)
	err := cursor.All(ctx, &data)
	if err != nil {
		log.Println("Error: ", err.Error())
		return data, err
	}
	return data, nil
}

/*Store Store a tweet in the database */
func (Model *Tweet) Store() (bool, error) {
	var Tweets = helpers.EstablishDriver("tweets")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	Model.CreatedAt = time.Now()
	if _, err := Tweets.InsertOne(ctx, Model); err != nil {
		return false, err
	}
	return true, nil
}

/*Update Update a tweet in the database */
func (Model *Tweet) Update(IDTweet string) (*Tweet, error) {
	var Tweets = helpers.EstablishDriver("tweets")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(IDTweet)
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
	var Tweets = helpers.EstablishDriver("tweets")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(IDTweet)
	filter := bson.M{"_id": bson.M{"$eq": ID}}
	if result, err := Tweets.DeleteOne(ctx, filter); err != nil || result.DeletedCount < 1 {
		return errors.New("It is not possible to remove the resource")
	}
	return nil
}

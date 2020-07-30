package models

import (
	"context"
	"time"

	"github.com/Jose-Guerrero-Developer/twittorbackend/database"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Profile struct
type Profile struct {
	User
}

// Get get profile model
func (profile *Profile) Get(ID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(ID)

	db := database.Connection.Database("twittor")
	users := db.Collection("users")
	condition := bson.M{
		"_id": objID,
	}

	err := users.FindOne(ctx, condition).Decode(&profile.User)
	profile.Password = ""

	if err != nil {
		return err
	}

	return nil
}

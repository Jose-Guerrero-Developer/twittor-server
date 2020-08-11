package helpers

import (
	"context"
	"strconv"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/utils/request"

	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*GetCollection returns a collection from the database */
func (Helper *Driver) GetCollection() *mongo.Collection {
	var GalexDatabase database.Driver
	return GalexDatabase.Client().Database(GalexDatabase.GetDatabaseName()).Collection(Helper.collection)
}

/*Find Mongo Find */
func (Helper *Driver) Find(ctx context.Context, filter bson.M) (*mongo.Cursor, error) {
	var GalexRequest request.Driver
	opts := options.Find()
	opts.SetLimit(GalexRequest.GetCount())
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
	opts.SetSkip((GalexRequest.GetPage() - 1) * GalexRequest.GetCount())
	GalexRequest.AddHeader("X-Current-Page", strconv.FormatInt(GalexRequest.GetPage(), 10))
	return Helper.GetCollection().Find(ctx, filter, opts)
}

/*FindOne Mongo FindOne */
func (Helper *Driver) FindOne(ctx context.Context, filter bson.M) *mongo.SingleResult {
	return Helper.GetCollection().FindOne(ctx, filter)
}

/*Aggregate Mongo Aggregate */
func (Helper *Driver) Aggregate(ctx context.Context, pepiline []bson.M) (*mongo.Cursor, error) {
	return Helper.GetCollection().Aggregate(ctx, pepiline)
}

/*InsertOne Mongo InsertOne */
func (Helper *Driver) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	return Helper.GetCollection().InsertOne(ctx, document)
}

/*UpdateOne Mongo UpdateOne */
func (Helper *Driver) UpdateOne(ctx context.Context, filter bson.M, update bson.M) (*mongo.UpdateResult, error) {
	return Helper.GetCollection().UpdateOne(ctx, filter, update)
}

/*DeleteOne Mongo DeleteOne */
func (Helper *Driver) DeleteOne(ctx context.Context, filter bson.M) (*mongo.DeleteResult, error) {
	return Helper.GetCollection().DeleteOne(ctx, filter)
}

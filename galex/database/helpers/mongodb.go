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

/*Collection returns a collection from the database */
func (Helper *Driver) Collection(name string) *mongo.Collection {
	var GalexDatabase database.Driver
	return GalexDatabase.Client().Database(GalexDatabase.GetDatabaseName()).Collection(name)
}

/*Find Mongo Find */
func (Helper *Driver) Find(ctx context.Context, table string, filter bson.M) (*mongo.Cursor, error) {
	var GalexRequest request.Driver
	var GalexDatabase database.Driver

	collection := GalexDatabase.Client().Database(GalexDatabase.GetDatabaseName()).Collection(table)
	opts := options.Find()
	opts.SetLimit(GalexRequest.GetCount())
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
	opts.SetSkip((GalexRequest.GetPage() - 1) * GalexRequest.GetCount())
	GalexRequest.AddHeader("X-Current-Page", strconv.FormatInt(GalexRequest.GetPage(), 10))
	return collection.Find(ctx, filter, opts)
}

/*FindOne Mongo FindOne */
func (Helper *Driver) FindOne(ctx context.Context, table string, filter bson.M) *mongo.SingleResult {
	var GalexDatabase database.Driver

	collection := GalexDatabase.Client().Database(GalexDatabase.GetDatabaseName()).Collection(table)
	return collection.FindOne(ctx, filter)
}

/*InsertOne Mongo InsertOne */
func (Helper *Driver) InsertOne(ctx context.Context, table string, document interface{}) (*mongo.InsertOneResult, error) {
	collection := Helper.Collection(table)
	return collection.InsertOne(ctx, document)
}

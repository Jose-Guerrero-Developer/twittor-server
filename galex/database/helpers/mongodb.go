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

/*Find Mongo find */
func (Helper *Driver) Find(ctx context.Context, table string, filter bson.M) (*mongo.Cursor, error) {
	var GalexRequest request.Driver
	var GalexDatabase database.Driver

	collections := GalexDatabase.Client().Database(GalexDatabase.GetDatabaseName()).Collection(table)
	opts := options.Find()
	opts.SetLimit(GalexRequest.GetCount())
	opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
	opts.SetSkip((GalexRequest.GetPage() - 1) * GalexRequest.GetCount())
	GalexRequest.AddHeader("X-Current-Page", strconv.FormatInt(GalexRequest.GetPage(), 10))
	return collections.Find(ctx, filter, opts)
}

/*Collection returns a collection from the database */
func (Helper *Driver) Collection(name string) *mongo.Collection {
	var GalexDatabase database.Driver
	return GalexDatabase.Client().Database(GalexDatabase.GetDatabaseName()).Collection(name)
}

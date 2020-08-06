package helpers

import (
	"github.com/Jose-Guerrero-Developer/twittorbackend/galex/database"
	"go.mongodb.org/mongo-driver/mongo"
)

/*Collection returns a collection from the database */
func (Helper *Driver) Collection(name string) *mongo.Collection {
	var GalexDatabase database.Driver
	return GalexDatabase.Client().Database(GalexDatabase.GetDatabaseName()).Collection(name)
}

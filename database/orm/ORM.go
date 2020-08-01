package orm

import (
	"github.com/Jose-Guerrero-Developer/twittorbackend/database"
	"go.mongodb.org/mongo-driver/mongo"
)

/*_Database driver database instance */
var _Database database.Driver

/*Driver structure manages ORM */
type Driver struct{}

/*Collection returns a collection from the database */
func (Controller *Driver) Collection(name string) *mongo.Collection {
	return _Database.Client().Database(_Database.GetDatabaseName()).Collection(name)
}

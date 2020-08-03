package helpers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

/*Driver structure manages ORM */
type Driver struct{}

/*Collection returns a collection from the database */
func (Controller *Driver) Collection(name string) *mongo.Collection {
	return _Database.Client().Database(_Database.GetDatabaseName()).Collection(name)
}

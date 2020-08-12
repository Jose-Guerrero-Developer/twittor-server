package database

import (
	"context"
	"time"

	"github.com/devJGuerrero/twittor-server/galex/configuration"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/*Driver Source of the package */
type Driver struct {
	databaseName string
}

/*storage Store package status */
var storage *Driver

/*Client mongo client instance */
var Client *mongo.Client

/*EstablishDriver Set package driver */
func (Database *Driver) EstablishDriver(name string) error {
	var err error
	storage = new(Driver)
	Client, err = createClientDatabase()
	if err != nil {
		return err
	}
	var GalexConfigs configuration.Driver
	storage.SetDatabaseName(GalexConfigs.Get("DATABASE_NAME"))
	return nil
}

/*Client return mongo client instance */
func (Database *Driver) Client() *mongo.Client {
	return Client
}

/*GetStatus check the connection to the database */
func (Database *Driver) GetStatus() (bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	err := Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return false, err.Error()
	}
	return true, string("")
}

/*SetDatabaseName set name database */
func (Database *Driver) SetDatabaseName(name string) {
	storage.databaseName = name
}

/*GetDatabaseName return database name */
func (Database *Driver) GetDatabaseName() string {
	return storage.databaseName
}

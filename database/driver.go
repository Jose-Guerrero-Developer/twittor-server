package database

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/*_Context stores the status of the database driver */
var _Context *Driver

/*Client mongo client instance */
var Client *mongo.Client

/*Driver structure manages database */
type Driver struct {
	databaseName string
}

/*LoadDriver establishes the connection to the database routines */
func (Controller *Driver) LoadDriver(databaseName string) error {
	_Context = new(Driver)
	MongoClient, err := _CreateClientDatabase()
	if err != nil {
		return err
	}
	Client = MongoClient
	Controller.SetDatabaseName(databaseName)
	if err := recover(); err != nil {
		return errors.New("An error occurred in setting up the database driver")
	}
	return nil
}

/*Client return mongo client instance */
func (Controller *Driver) Client() *mongo.Client {
	return Client
}

/*GetStatus check the connection to the database */
func (Controller *Driver) GetStatus() (bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	err := Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return false, err.Error()
	}
	return true, string("")
}

/*SetDatabaseName Set up the database */
func (Controller *Driver) SetDatabaseName(name string) {
	_Context.databaseName = name
}

/*GetDatabaseName return database name */
func (Controller *Driver) GetDatabaseName() string {
	return _Context.databaseName
}

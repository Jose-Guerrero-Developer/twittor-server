package database

import (
	"context"
	"log"
	"time"

	"github.com/devJGuerrero/twittor-server/galex/configuration"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/*createClientDatabase create the instance with the mongo client */
func createClientDatabase() (*mongo.Client, error) {
	ctxConnection, cancelConnection := context.WithTimeout(context.Background(), 10*time.Minute)
	ctxCheckConnection, cancelCheckConnection := context.WithTimeout(context.Background(), 2*time.Minute)

	defer func() {
		cancelConnection()
		cancelCheckConnection()
	}()

	var GalexConfigs configuration.Driver
	MongoClient, err := mongo.Connect(ctxConnection, options.Client().ApplyURI("mongodb+srv://"+GalexConfigs.Get("DATABASE_USERNAME")+":"+GalexConfigs.Get("DATABASE_PASSWORD")+"@"+GalexConfigs.Get("DATABASE_HOST")+".r04jg.mongodb.net/<dbname>?retryWrites=true&w=majority"))
	if err != nil {
		return MongoClient, err
	}

	err = MongoClient.Ping(ctxCheckConnection, readpref.Primary())
	if err != nil {
		return MongoClient, err
	}
	log.Println("Conexión establecidad al servidor de base de datos")
	return MongoClient, nil
}

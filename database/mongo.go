package database

import (
	"context"
	"log"
	"time"

	"github.com/Jose-Guerrero-Developer/twittorbackend/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/*_CreateClientDatabase create the instance with the mongo client */
func _CreateClientDatabase() (*mongo.Client, error) {
	ctxConnection, cancelConnection := context.WithTimeout(context.Background(), 10*time.Minute)
	ctxCheckConnection, cancelCheckConnection := context.WithTimeout(context.Background(), 2*time.Minute)

	defer func() {
		cancelConnection()
		cancelCheckConnection()
	}()

	var Configs configs.Driver
	MongoClient, err := mongo.Connect(ctxConnection, options.Client().ApplyURI("mongodb+srv://"+Configs.Get("DATABASE_USERNAME")+":"+Configs.Get("DATABASE_PASSWORD")+"@"+Configs.Get("DATABASE_NAME")+".r04jg.mongodb.net/<dbname>?retryWrites=true&w=majority"))
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

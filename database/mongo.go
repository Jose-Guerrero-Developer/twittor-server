package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connection mongo
var Connection = connection()

func connection() (client *mongo.Client) {
	ctxConnection, cancelConnection := context.WithTimeout(context.Background(), 10*time.Second)
	ctxCheckConnection, cancelCheckConnection := context.WithTimeout(context.Background(), 2*time.Second)

	client, err := mongo.Connect(ctxConnection, options.Client().ApplyURI("mongodb+srv://root:root@estudios.r04jg.mongodb.net/<dbname>?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	err = client.Ping(ctxCheckConnection, readpref.Primary())
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	log.Println("Conexión establecidad al servidor de base de datos")

	defer func() {
		cancelConnection()
		cancelCheckConnection()
	}()
	return
}

// CheckConnectionStatus Check connection status
func CheckConnectionStatus() (status bool, error string) {
	status = true
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := Connection.Ping(ctx, readpref.Primary())

	if err != nil {
		status = false
		error = err.Error()
	}

	return
}

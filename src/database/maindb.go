package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MainDB struct {
	Client *mongo.Client
}

func (db *MainDB) Connect(mongodburi string) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongodburi).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		panic(fmt.Errorf("error connecting to database: %w", err))
	}

	db.Client = client
}

func (db *MainDB) Disconnect() {
	if err := db.Client.Disconnect(context.TODO()); err != nil {
		panic(fmt.Errorf("error disconnecting from database: %w", err))
	}
}

func (db *MainDB) Ping(databaseName string) bool {
	if err := db.Client.Database(databaseName).RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		panic(fmt.Errorf("error pinging database: %w", err))
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return true
}

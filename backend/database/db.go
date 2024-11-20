package database

import (
	"context"
	"fmt"
	"musicboxd/hlp"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	instance *DB
	once     sync.Once
)

type DB struct {
	database *mongo.Database
	client   *mongo.Client
}

func GetDB() *DB {
	once.Do(func() {
		instance = ConnectDB()
	})
	return instance
}

func ConnectDB() *DB {
	uri := hlp.Envs["MONGO_URI"]
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB")
	db := client.Database(hlp.Envs["MONGO_DB"])

	return &DB{
		database: db,
		client:   client,
	}
}

func (db *DB) GetCollection(name string) *mongo.Collection {
	return db.database.Collection(name)
}

func (db *DB) Close() error {
	return db.client.Disconnect(context.TODO())
}

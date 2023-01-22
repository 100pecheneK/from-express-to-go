package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var clientInstanceError error

var mongoOnce sync.Once

type Collection string

const (
	ProductsCollection Collection = "products"
)

const (
	uri      = "mongodb://mongo:27017"
	Database = "products-api"
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(uri)
		clientOptions.SetAuth(options.Credential{
			Username: "admin",
			Password: "password",
		})
		client, err := mongo.Connect(context.TODO(), clientOptions)
		clientInstance = client
		clientInstanceError = err
	})
	return clientInstance, clientInstanceError
}

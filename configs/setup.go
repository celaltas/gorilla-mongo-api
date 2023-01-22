package configs

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var ctx = context.TODO()

func ConnectDB() *mongo.Client {

	uri, err := EnvMongoUri()
	if err != nil {
		fmt.Println(err)
	}
	options := options.Client().ApplyURI(uri).SetTimeout(3 * time.Second)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client

}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}

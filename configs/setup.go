package configs

import (
	"context"
	"gin-mongo-api/logger"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	database = "testdb"
)

var Client *mongo.Client

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		logger.ErrorLogger.Fatalln(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logger.ErrorLogger.Fatalln(err)
	}

	// ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		logger.ErrorLogger.Fatalln(err)
	}

	logger.InfoLogger.Println("Connected to MongoDB!")
	return client
}

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(database).Collection(collectionName)
	return collection
}

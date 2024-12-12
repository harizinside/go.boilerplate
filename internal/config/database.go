package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client   *mongo.Client
	database *mongo.Database
)

type MongoConfig struct {
	URI      string
	Database string
}

func ConnectMongo(DATABASE_URL string, DATABASE_NAME string) error {
	cfg := MongoConfig{
		URI:      DATABASE_URL,
		Database: DATABASE_NAME,
	}

	clientOptions := options.Client().ApplyURI(cfg.URI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return err
	}

	database = client.Database(cfg.Database)
	log.Println("Connected to MongoDB")

	return nil
}

func DisconnectMongo() {
	if client == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		log.Printf("Failed to disconnect MongoDB: %v", err)
	} else {
		log.Println("Disconnected from MongoDB")
	}
}

func GetDatabase() *mongo.Database {
	return database
}

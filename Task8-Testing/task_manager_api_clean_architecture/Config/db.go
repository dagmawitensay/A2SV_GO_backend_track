package config

import (
	"context"
	"log"
	domain "task_manager_api_clean_architecture/Domain"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func GetDB(config *domain.Config) (*mongo.Database) {
	client, err := connectToMongoDB(config.Database.DBURI)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(config.Database.DbName)

	return db
}

func connectToMongoDB(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")
	return client, nil
}
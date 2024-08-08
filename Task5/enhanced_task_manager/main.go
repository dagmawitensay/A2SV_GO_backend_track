package main

import (
	"context"
	"enhanced_task_manager/data"
	"enhanced_task_manager/router"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	collection, err := getCollection()

	if err != nil {
		log.Fatal(err)
	}

	taskService := data.NewTaskService(collection)
	r := router.SetupRouter(*taskService)
	r.Run(":8080")
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

func getCollection() (*mongo.Collection, error) {
	client, err := connectToMongoDB("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	collection := client.Database("taskdb").Collection("tasks")

	return collection, nil
}
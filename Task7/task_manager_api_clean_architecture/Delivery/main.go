package main

import (
	"context"
	"log"
	"task_manager_api_clean_architecture/Delivery/routers"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	db, err := getDBConnection()

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	
	routers.SetupRouter(10 * time.Second, db, r)

	r.Run(":8080")
}

func getDBConnection() (*mongo.Database, error) {
	client, err := connectToMongoDB("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db := client.Database("taskdb-cleanArc")

	return db, nil
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
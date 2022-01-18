package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sudipto-003/sweet-gin/repository"
	"github.com/sudipto-003/sweet-gin/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	mongUser := os.Getenv("root")
	mongoPass := os.Getenv("pass")

	uri := fmt.Sprintf("mongodb://%s:%s@0.0.0.0:27017", mongUser, mongoPass)
	mongoClient, err := repository.CreateMongoConnection(context.Background(), uri)
	if err != nil {
		log.Panic(err)
	}
	defer mongoClient.Disconnect(context.Background())

	// collection := mongoClient.Database("sweetgin").Collection("parcel")
	repo := &repository.Repos{MongoClient: mongoClient}

	r := router.GetHttpRouter(repo)

	r.Run(":8080")
}

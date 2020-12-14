package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Connector = ConnectCubes()
func ConnectCubes() *mongo.Collection {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	fmt.Println("Server is starting...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URL"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("firstApi")
	cubeCollection := database.Collection("Cube")

	fmt.Println("Connected to MongoDB!")
	return cubeCollection
}
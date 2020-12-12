package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectCubes() *mongo.Collection {
	fmt.Println("Server is starting...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//clientOptions := options.Client().ApplyURI("mongodb+srv://tuanconbu:hieu1234@go-api.rjhwq.mongodb.net/firstApi?retryWrites=true&w=majority")
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
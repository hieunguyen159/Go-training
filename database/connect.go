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

var Connector = ConnectCubes()
var ConnectorEmails = ConnectEmails()
var ConnectorUsers = ConnectUsers()

func ConnectCubes() *mongo.Collection {
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
func ConnectEmails() *mongo.Collection {
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
	emailsCollection := database.Collection("Emails")

	fmt.Println("Connected to MongoDB!")
	return emailsCollection
}
func ConnectUsers() *mongo.Collection {
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
	usersCollection := database.Collection("Users")

	fmt.Println("Connected to MongoDB!")
	return usersCollection
}

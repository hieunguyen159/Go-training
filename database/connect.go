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

// cubes collection
var Connector = ConnectCubes()

// emails collection
var ConnectorEmails = ConnectEmails()

// users collection
var ConnectorUsers = ConnectUsers()

func ConnectCubes() *mongo.Collection {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("Error loading .env file")
	}
	fmt.Println("Server is starting...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URL"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Println(err)
	}
	database := client.Database("firstApi")
	cubeCollection := database.Collection("Cube")

	fmt.Println("Connected to MongoDB, collection: Cubes")
	return cubeCollection
}
func ConnectEmails() *mongo.Collection {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("Error loading .env file")
	}
	fmt.Println("Server is starting...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URL"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Println(err)
	}
	database := client.Database("firstApi")
	emailsCollection := database.Collection("Emails")

	fmt.Println("Connected to MongoDB, collection: Emails")
	return emailsCollection
}
func ConnectUsers() *mongo.Collection {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("Error loading .env file")
	}
	fmt.Println("Server is starting...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URL"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Println(err)
	}
	database := client.Database("firstApi")
	usersCollection := database.Collection("Users")

	fmt.Println("Connected to MongoDB, collection: Users")
	return usersCollection
}

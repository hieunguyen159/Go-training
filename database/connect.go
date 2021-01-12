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

var (
	MgoDatabase     *mongo.Database
	ConnectorEmails = ConnectEmails() // emails collection
	ConnectorUsers  = ConnectUsers()  // users collection
)

type DBIInterface interface {
	GetDatabase() *mongo.Collection
}

func NewDBI() DBIInterface {
	return &MongoDB{}
}
func NewDB(connectstring string, dbName string) {
	connection(connectstring, dbName)
}
func connection(connectstring string, dbName string) {
	fmt.Println("Server is starting...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(connectstring)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Driver offical - connected to MongoDB!")
	MgoDatabase = client.Database(dbName)
	fmt.Println("Using DB:", dbName)
}

func NewCollectionDB(collection string) *MongoDB {
	return &MongoDB{c: MgoDatabase.Collection(collection)}
}
func (dbi *MongoDB) GetDatabase() *mongo.Collection {
	return dbi.c
}

type MongoDB struct {
	c *mongo.Collection
}

type DataService struct {
	CubesCollection  *mongo.Collection
	EmailsCollection *mongo.Collection
	UsersCollection  *mongo.Collection
}

func LoadConfig() (*DataService, error) {
	cubesCollection := NewCollectionDB("Cube")
	emailsCollection := NewCollectionDB("Emails")
	usersCollection := NewCollectionDB("Users")
	return &DataService{
		CubesCollection:  cubesCollection.c,
		EmailsCollection: emailsCollection.c,
		UsersCollection:  usersCollection.c,
	}, nil
}

func ConnectEmails() *mongo.Collection {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Printf("Error loading .env file")
	// }
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
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Printf("Error loading .env file")
	// }
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

package main

import (
	"context"
	"log"

	// "os"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
type Cube struct {
	Currency  string `bson:"currency"`
	Rate string `bson:"rate"`
}
type Cubes struct{
	ID        primitive.ObjectID           `bson:"_id"`
	Time      string                       `bson:"time"`
	Cubes 	[]Cube 				    `bson:"Cube"`
}
func main() {

	router := gin.Default()

	router.GET("/newest-rates",newestRateFunc)

	router.Run(":8080")
}


func newestRateFunc(c *gin.Context){
	cubeCollection := db.ConnectCubes();
	var Cubes []Cubes
	data,_ := cubeCollection.Find(context.Background(),bson.M{})
	defer data.Close(context.Background())
	error := data.All(context.Background(),&Cubes)
	if error != nil {
		log.Fatal(error)
	}
	c.JSON(http.StatusOK, Cubes)
}
package controllers

import (
	db "api/database"
	models "api/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllEmails(c *gin.Context) {
	emailCollection := db.ConnectorEmails
	var emails []models.Email
	data, _ := emailCollection.Find(context.Background(), bson.M{})
	defer data.Close(context.Background())
	err := data.All(context.Background(), &emails)
	if err == nil {
		c.JSON(http.StatusOK, emails)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error": "Can't query database",
		})
	}
}

func TurnOffRemindEmail(c *gin.Context) {
	turnOffID := c.Param("id")
	emailCollection := db.ConnectorEmails
	oid, _ := primitive.ObjectIDFromHex(turnOffID)
	filter := bson.D{{"_id", oid}}
	update := bson.D{{"$set", bson.D{{"reminded", false}}}}
	_, err := emailCollection.UpdateOne(context.Background(), filter, update)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "Turn Off OK",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "Error",
		})
	}
}

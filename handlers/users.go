package handlers

import (
	db "api/database"
	"api/models"
	jwt "api/utils"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllUsers(c *gin.Context) {
	email, err := jwt.ExtractEmailFromToken(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error"})
		return
	}
	usersCollection := db.ConnectorUsers
	var users []models.User
	data, err := usersCollection.Find(context.Background(), bson.M{"email": bson.M{"$ne": email}})
	defer data.Close(context.Background())
	data.All(context.Background(), &users)
	c.JSON(200, users)
}
func ToggleUser(c *gin.Context) {
	roles, err := jwt.ExtractRolesFromToken(c)
	fmt.Println("roles", roles)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error"})
		return
	}
	if roles == "ADMIN" {
		var status models.Status
		usersCollection := db.ConnectorUsers
		toggleID := c.Param("id")

		c.ShouldBindBodyWith(&status, binding.JSON)
		oid, _ := primitive.ObjectIDFromHex(toggleID)
		_, err := usersCollection.UpdateOne(context.Background(), bson.M{"_id": oid}, bson.M{"$set": bson.M{"active": status.Status}})
		if err == nil {
			c.JSON(200, gin.H{
				"message": "Success",
			})
			return
		} else {
			c.JSON(500, gin.H{
				"message": "Internal Server Error"})
			return
		}
	} else {
		c.JSON(403, gin.H{
			"message": "Not permission to access this resource"})
		return
	}
}

func SetRolesUser(c *gin.Context) {
	roles, err := jwt.ExtractRolesFromToken(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error"})
		return
	}
	if roles == "ADMIN" {
		var rolesChange models.Role
		usersCollection := db.ConnectorUsers
		userID := c.Param("id")

		c.ShouldBindBodyWith(&rolesChange, binding.JSON)
		if len(rolesChange.Roles) > 0 {
			oid, _ := primitive.ObjectIDFromHex(userID)
			_, err := usersCollection.UpdateOne(context.Background(), bson.M{"_id": oid}, bson.M{"$set": bson.M{"roles": rolesChange.Roles}})
			if err == nil {
				c.JSON(200, gin.H{
					"message": "Success",
				})
				return
			} else {
				c.JSON(500, gin.H{
					"message": "Internal Server Error"})
				return
			}
		} else {
			c.JSON(400, gin.H{
				"message": "Roles are not empty"})
			return
		}

	} else {
		c.JSON(403, gin.H{
			"message": "Not permission to access this resource"})
		return
	}
}

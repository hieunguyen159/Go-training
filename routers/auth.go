package routers

import (
	db "api/database"
	helpers "api/helpers"
	models "api/models"
	jwt "api/utils"
	"context"
	"fmt"
	"log"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(c *gin.Context) {
	var loggedInUser models.LoggedInUser
	if err := c.ShouldBindBodyWith(&loggedInUser, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}
	if govalidator.IsNull(loggedInUser.Email) || govalidator.IsNull(loggedInUser.Password) {
		c.JSON(400, gin.H{"message": "Data can not empty"})
		return
	}

	loggedInUser.Email = models.Santize(loggedInUser.Email)
	loggedInUser.Password = models.Santize(loggedInUser.Password)

	usersCollection := db.ConnectorUsers

	var result models.User
	err := usersCollection.FindOne(context.TODO(), bson.M{"email": loggedInUser.Email}).Decode(&result)
	if err != nil {
		c.JSON(400, gin.H{"message": "Email or Password incorrect"})
		return
	}
	hashedPassword := fmt.Sprintf("%v", result.Password)
	err = models.CheckPasswordHash(hashedPassword, loggedInUser.Password)
	if err != nil {
		c.JSON(401, gin.H{"message": "Email or Password incorrect"})
		return
	}
	if result.Active {
		token, errCreate := jwt.Create(loggedInUser.Email)

		if errCreate != nil {
			c.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}

		if helpers.CheckExist("ADMIN", result.Roles) {
			c.JSON(200, gin.H{
				"token": token,
			})
		} else {
			c.JSON(401, gin.H{"message": "No permission"})
		}
	} else {
		c.JSON(403, gin.H{
			"message": "This account is banned",
		})
	}

}
func Register(c *gin.Context) {
	var signInUser models.User
	if err := c.ShouldBindBodyWith(&signInUser, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}
	if govalidator.IsNull(signInUser.Email) || govalidator.IsNull(signInUser.Password) || len(signInUser.Roles) == 0 {
		c.JSON(400, gin.H{"message": "Bad request"})
		return
	}
	if !govalidator.IsEmail(signInUser.Email) {
		c.JSON(400, gin.H{"message": "Email is invalid"})
		return
	}
	signInUser.Email = models.Santize(signInUser.Email)
	signInUser.Password = models.Santize(signInUser.Password)

	usersCollection := db.ConnectorUsers

	var result bson.M
	errFindEmail := usersCollection.FindOne(context.Background(), bson.M{"email": signInUser.Email}).Decode(&result)
	if errFindEmail == nil {
		c.JSON(409, gin.H{"message": "User does exists"})
		return
	}
	password, err := models.Hash(signInUser.Password)

	if err != nil {
		c.JSON(500, gin.H{"message": "Register has failed"})
		return
	}
	signInUser.Password = password
	newUser := models.NewUserBson(signInUser)
	_, errs := usersCollection.InsertOne(context.Background(), newUser)
	if errs != nil {
		c.JSON(500, gin.H{"message": "Register has failed"})
		return
	}
	emailCollection := db.ConnectorEmails
	newAlertEmail := models.NewEmailBSon(signInUser.Email)
	emailCollection.InsertOne(context.Background(), newAlertEmail)
	c.JSON(201, gin.H{"message": "Register Succesfully"})
}

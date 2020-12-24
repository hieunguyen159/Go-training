package controllers

import (
	db "api/database"
	helpers "api/helpers"
	models "api/models"
	"context"
	"crypto/tls"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	gomail "gopkg.in/mail.v2"
)

func SendToAllUser(c *gin.Context) {
	emailCollection := db.ConnectorEmails
	var form models.Form
	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}
	latestCubes := GetNewest(c)
	m := gomail.NewMessage()
	var email models.Email

	recipients := form.Receiver

	for _, receiver := range recipients {
		email.Email = receiver
		email.Status = "waiting"
		res, _ := emailCollection.InsertOne(context.Background(), email)
		id := res.InsertedID
		go SendMail(emailCollection, form, receiver, latestCubes, m, id)

		var userEmailSent []models.Email
		emails, _ := emailCollection.Find(context.Background(), bson.M{})
		defer emails.Close(context.Background())
		emails.All(context.Background(), &userEmailSent)
		allEmails := make([]models.Email, 0)
		for _, user := range userEmailSent {
			allEmails = append(allEmails, user)
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
func SendMail(emailCollection *mongo.Collection, form models.Form, receiver string, latestCubes models.DateCube, m *gomail.Message, id interface{}) {
	m.SetHeader("From", form.Email)
	m.SetHeader("To", receiver)
	m.SetHeader("Subject", "Ahehehehe")
	m.SetBody("text/plain", helpers.CreateKeyValuePairs(latestCubes.Rates))

	d := gomail.Dialer{Host: "mailhog", Port: 1025}

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	emailCollection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"status": "done"}})
}

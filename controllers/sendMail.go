package controllers

import (
	db "api/database"
	helpers "api/helpers"
	models "api/models"
	socket "api/websocket"
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
	latestCubes := NewestRates()
	emailCollection := db.ConnectorEmails
	emailCollection.Drop(context.TODO())

	var form models.Form
	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}

	m := gomail.NewMessage()

	recipients := form.Receiver
	allEmails := make([]models.Email, 0)
	socketEmails := make([]models.Email, len(recipients))
	for i, receiver := range recipients {
		email := models.NewEmailBSon(receiver)
		log.Println(email.ID)
		res, _ := emailCollection.InsertOne(context.Background(), email)
		id := res.InsertedID
		allEmails = append(allEmails, email)
		go SendMail(email, emailCollection, form, receiver, latestCubes, m, id, i, socketEmails)
	}
	c.JSON(http.StatusOK, allEmails)

}
func SendMail(email models.Email, emailCollection *mongo.Collection, form models.Form, receiver string, latestCubes models.DateCube, m *gomail.Message, id interface{}, i int, socketEmails []models.Email) {
	m.SetHeader("From", form.Email)
	m.SetHeader("To", receiver)
	m.SetHeader("Subject", "Ahehehehe")
	m.SetBody("text/plain", helpers.CreateKeyValuePairs(latestCubes.Rates))

	d := gomail.Dialer{Host: "mailhog", Port: 1025}

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	_, err := emailCollection.UpdateOne(context.Background(), bson.M{"_id": id}, bson.M{"$set": bson.M{"status": "done"}})
	if err == nil {
		email.Status = "done"
		socketEmails[i] = email
		socket.PushMessage(socketEmails)
	}

}

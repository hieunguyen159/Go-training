package controllers

import (
	db "api/database"
	helpers "api/helpers"
	models "api/models"
	socket "api/websocket"
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	gomail "gopkg.in/mail.v2"
)

func SendToAllUser(c *gin.Context) {
	latestCubes := NewestRates()
	emailCollection := db.ConnectorEmails
	// emailCollection.Drop(context.TODO())

	var form models.Form
	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}

	m := gomail.NewMessage()
	var existedEmails []models.Email

	data, _ := emailCollection.Find(context.Background(), bson.M{})
	defer data.Close(context.Background())
	data.All(context.Background(), &existedEmails)

	recipientsReq := form.Receiver
	var recipients []string
	// check exists email in DB
	if len(existedEmails) > 0 {
		for _, mailReq := range recipientsReq {
			for _, mailDB := range existedEmails {
				if mailReq != mailDB.Email {
					recipients = append(recipients, mailReq)
				}
			}
		}
		recipientsReq = recipients
	}
	finalRecipients := helpers.RemoveDuplicate(recipientsReq)
	log.Println("recipients", finalRecipients)
	allEmails := make([]models.Email, 0)
	socketEmails := make([]models.Email, len(finalRecipients))
	if len(finalRecipients) > 0 {
		for i, receiver := range finalRecipients {
			email := models.NewEmailBSon(receiver)
			res, _ := emailCollection.InsertOne(context.Background(), email)
			id := res.InsertedID
			allEmails = append(allEmails, email)
			go SendMail(email, emailCollection, form, receiver, latestCubes, m, id, i, socketEmails)
		}
		c.JSON(http.StatusOK, allEmails)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "Sent mails failed",
		})
	}
}
func SendMail(email models.Email, emailCollection *mongo.Collection, form models.Form, receiver string, latestCubes models.DateCube, m *gomail.Message, id interface{}, i int, socketEmails []models.Email) {
	if form.Email != "" {
		m.SetHeader("From", form.Email)
		m.SetHeader("To", receiver)
		m.SetHeader("Subject", "Ahehehehe")
		stringID := id.(primitive.ObjectID).Hex()
		bodyText := fmt.Sprintf(helpers.CreateKeyValuePairs(latestCubes.Rates)+"Redirect to http://localhost:3000/%s to turn off the alert everyday", stringID)
		m.SetBody("text/plain", bodyText)

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
	} else {
		return
	}

}

func SendMailEveryday() {
	godotenv.Load()
	var form models.Form
	var receivers []models.Email
	var emailsReceivers []string
	emailCollection := db.ConnectorEmails
	latestCubes := NewestRates()
	m := gomail.NewMessage()

	filter := bson.M{"reminded": true}
	data, err := emailCollection.Find(context.Background(), filter)
	defer data.Close(context.Background())
	data.All(context.Background(), &receivers)
	for _, email := range receivers {
		emailsReceivers = append(emailsReceivers, email.Email)
	}
	socketEmails := make([]models.Email, len(emailsReceivers))
	if err == nil {
		form.Email = os.Getenv("MAIL")
		form.Receiver = emailsReceivers
	}
	for i, receiver := range emailsReceivers {
		email := models.NewEmailBSon(receiver)
		res, _ := emailCollection.InsertOne(context.Background(), email)
		id := res.InsertedID
		go SendMail(email, emailCollection, form, receiver, latestCubes, m, id, i, socketEmails)
	}

}

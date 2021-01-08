package handlers

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
	m := gomail.NewMessage()

	var form models.Form
	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}
	requestRecipients := form.Receiver

	var emailsCollectionInDB []models.Email
	data, err := emailCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return
	}
	defer data.Close(context.Background())
	data.All(context.Background(), &emailsCollectionInDB)

	var emailsInDB []string
	for _, item := range emailsCollectionInDB {
		emailsInDB = append(emailsInDB, item.Email)
	}
	var recipients []string

	// check exists email in DB
	for _, email := range requestRecipients {
		_, found := helpers.Find(email, emailsInDB)
		if !found {
			recipients = append(recipients, email)
		}
	}

	log.Println("recipients", recipients)
	allEmails := make([]models.Email, 0)
	socketEmails := make([]models.Email, len(recipients))
	if len(recipients) > 0 {
		for i, receiver := range recipients {
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
		bodyText := fmt.Sprintf(helpers.CreateKeyValuePairs(latestCubes.Rates)+"Redirect to http://localhost:3000/alert/%s to turn off the alert everyday", stringID)
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

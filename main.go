package main

import (
	ctrl "api/controllers"
	db "api/database"
	"api/models"
	"context"
	"encoding/xml"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/robfig/cron.v2"
)

func main() {

	router := gin.Default()

	router.GET("/newest-rates", ctrl.GetNewestRates)
	router.POST("/random-rates", ctrl.GetRandomRates)
	router.GET("/value-per-currency", ctrl.GetPropertyOfAll)
	GetXMLfile()
	router.Run(":8080")
	c := cron.New()
	c.AddFunc("@hourly", GetXMLfile)
	c.Start()
	c.Stop()
}

func GetXMLfile() {
	resp, err := http.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	var envelope models.Envelope
	if err = xml.NewDecoder(resp.Body).Decode(&envelope); err != nil {
		log.Println(err)
	}

	var ui []interface{}
	for _, t := range envelope.Envelope.BigCube {
		if t.Time == "2020-12-04" {
			ui = append(ui, t)
		}

	}
	db.Connector.InsertMany(context.Background(), ui)
}

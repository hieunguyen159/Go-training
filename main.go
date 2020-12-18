package main

import (
	ctrl "api/controllers"
	db "api/database"
	"api/models"
	"context"
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	
	router.GET("/newest-rates",ctrl.GetNewestRates)
	router.POST("/random-rates",ctrl.GetRandomRates)
	router.GET("/value-per-currency",ctrl.GetPropertyOfAll)
	GetXMLfile()
	router.Run(":8081")
}

func GetXMLfile() {
	resp,err := http.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml")
	//log.Println(resp)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	var envelope models.Envelope
	if err = xml.NewDecoder(resp.Body).Decode(&envelope); err != nil {
		log.Println(err)
	}

	var ui []interface{}
	for _, t := range envelope.Envelope.BigCube{
		ui = append(ui, t)
	}
	db.Connector.InsertMany(context.Background(),ui)
}
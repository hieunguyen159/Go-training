package main

import (
	ctrl "api/controllers"
	db "api/database"

	"github.com/gin-gonic/gin"
	"gopkg.in/robfig/cron.v2"
)

func main() {

	router := gin.Default()

	router.GET("/newest-rates", ctrl.GetNewestRates)
	router.POST("/random-rates", ctrl.GetRandomRates)
	router.GET("/value-per-currency", ctrl.GetPropertyOfAll)
	db.GetXMLfile()
	c := cron.New()
	c.Start()
	c.AddFunc("@every 0h0m1s", db.AddDataDaily)

	router.Run(":8081")
}

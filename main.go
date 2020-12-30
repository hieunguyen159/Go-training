package main

import (
	ctrl "api/controllers"
	db "api/database"
	socket "api/websocket"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/robfig/cron.v2"
)

func main() {

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/newest-rates", ctrl.GetNewestRates)
	router.POST("/random-rates", ctrl.GetRandomRates)
	router.GET("/value-per-currency", ctrl.GetPropertyOfAll)
	router.GET("/ws", socket.Echo)
	router.POST("/mail/send-all", ctrl.SendToAllUser)
	db.GetXMLfile()
	c := cron.New()
	c.Start()
	c.AddFunc("@daily", db.AddDataDaily)

	router.Run(":8080")
}

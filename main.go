package main

import (
	ctrl "api/controllers"
	db "api/database"
	socket "api/websocket"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/robfig/cron.v2"
)

func main() {
	godotenv.Load()
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/newest-rates", ctrl.GetNewestRates)
	router.POST("/random-rates", ctrl.GetRandomRates)
	router.GET("/value-per-currency", ctrl.GetPropertyOfAll)
	router.GET("/ws", socket.Echo)
	router.POST("/mail/send-all", ctrl.SendToAllUser)
	router.GET("/emails", ctrl.GetAllEmails)
	router.PUT("/emails/:id", ctrl.TurnOffRemindEmail)

	db.GetXMLfile()
	c := cron.New()
	c.Start()
	c.AddFunc("@daily", db.AddDataDaily)
	c.AddFunc("@daily", ctrl.SendMailEveryday)
	router.Run(":" + os.Getenv("PORT"))
}

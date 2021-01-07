package main

import (
	db "api/database"
	middlewares "api/middlewares"
	routers "api/routers"
	"os"
	"time"
	socket "api/websocket"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"
	"gopkg.in/robfig/cron.v2"
)

func main() {
	godotenv.Load()
	// Initialize a new Gin router
	router := gin.New()

	// Apply the middleware to the router (works with groups too)
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.GET("/newest-rates", middlewares.CheckJwt(routers.GetNewestRates))
	router.POST("/random-rates", routers.GetRandomRates)
	router.GET("/value-per-currency", routers.GetPropertyOfAll)

	router.GET("/ws", socket.Echo)

	router.POST("/mail/send-all", routers.SendToAllUser)
	router.GET("/emails", routers.GetAllEmails)
	router.PUT("/emails/:id", routers.TurnOffRemindEmail)

	router.POST("/auth/login", routers.Login)
	router.POST("/auth/register", routers.Register)

	router.GET("/users", middlewares.CheckJwt(routers.GetAllUsers))
	router.PUT("/users/roles/:id", middlewares.CheckJwt(routers.SetRolesUser))
	router.PUT("/users/active/:id", middlewares.CheckJwt(routers.ToggleUser))
	db.GetXMLfile()
	c := cron.New()
	c.Start()
	c.AddFunc("@daily", db.AddDataDaily)
	c.AddFunc("@daily", routers.SendMailEveryday)
	router.Run(":" + os.Getenv("PORT"))
}

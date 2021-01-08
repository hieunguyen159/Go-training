package main

import (
	db "api/database"
	handlers "api/handlers"
	routers "api/routers"
"log"
"github.com/joho/godotenv"
	"gopkg.in/robfig/cron.v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
	}

	db.GetXMLfile()
	c := cron.New()
	c.Start()
	c.AddFunc("@daily", db.AddDataDaily)
	c.AddFunc("@daily", handlers.SendMailEveryday)
	// Initialize a new Gin router
	routers.Router()
}

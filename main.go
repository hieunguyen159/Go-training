package main

import (
	db "api/database"
	handlers "api/handlers"
	routers "api/routers"

	"github.com/joho/godotenv"
	"gopkg.in/robfig/cron.v2"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	err = db.GetXMLfile()
	if err == nil {
		c := cron.New()
		c.Start()
		c.AddFunc("@daily", db.AddDataDaily)
		c.AddFunc("@daily", handlers.SendMailEveryday)
		// Initialize a new Gin router
		routers.Router()
	}
}

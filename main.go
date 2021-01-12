package main

import (
	db "api/database"
	handlers "api/handlers"
	routers "api/routers"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/robfig/cron.v2"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db.NewDB(os.Getenv("DB_URL"), os.Getenv("DB_NAME"))
	c := cron.New()
	c.Start()
	c.AddFunc("@daily", db.AddDataDaily)
	c.AddFunc("@daily", handlers.SendMailEveryday)

	dao, _ := db.LoadConfig()

	dbi := db.NewDBI()
	routers.Router(dbi, dao)

}

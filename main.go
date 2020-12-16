package main

import (
	ctrl "api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	router.GET("/newest-rates",ctrl.GetNewestRates)
	router.POST("/random-rates",ctrl.GetRandomRates)
	router.GET("/value-per-currency",ctrl.GetPropertyOfAll)

	router.Run(":8080")
}
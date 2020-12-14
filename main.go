package main

import (
	"context"
	"log"
	"strconv"

	// "os"

	db "api/database"
	models "api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)



func main() {
	router := gin.Default()
	
	router.GET("/newest-rates",func (c *gin.Context){
		cubeCollection := db.Connector
		var Cubes []models.Cubes
		var DateCubes models.DateCubes
		data,_ := cubeCollection.Find(context.Background(),bson.M{})
		defer data.Close(context.Background())
		error := data.All(context.Background(),&Cubes)
		if error != nil {
			log.Fatal(error)
		}
		rateResult := make(map[string]float64)
		for _,cubes := range Cubes {
			DateCubes.Date = cubes.Time
			for _,cube := range cubes.Cubes {
				s, _ := strconv.ParseFloat(cube.Rate, 64)
				rateResult[cube.Currency] = s
				DateCubes.Rates = rateResult
			}
		}
		c.JSON(http.StatusOK, DateCubes)
	})
	router.GET("/random-rates",func (c *gin.Context){
		cubeCollection := db.Connector
		var Cubes []models.Cubes
		var DateCubes models.DateCubes
		data,_ := cubeCollection.Find(context.Background(),bson.M{"time" : "2020-12-10"})
		defer data.Close(context.Background())
		error := data.All(context.Background(),&Cubes)
		if error != nil {
			log.Fatal(error)
		}
		rateResult := make(map[string]float64)
		for _,cubes := range Cubes {
			DateCubes.Date = cubes.Time
			for _,cube := range cubes.Cubes {
				s, _ := strconv.ParseFloat(cube.Rate, 64)
				rateResult[cube.Currency] = s
			
				DateCubes.Rates = rateResult
			}
		}
		c.JSON(http.StatusOK, DateCubes)
	})
	router.GET("/value-per-currency",func (c *gin.Context){
		cubeCollection := db.Connector
		var Cubes []models.Cubes
		var dateCubes models.DateCubes
		allcubes := make([]models.DateCubes,0)

		data,_ := cubeCollection.Find(context.Background(),bson.M{})
		defer data.Close(context.Background())
		error := data.All(context.Background(),&Cubes)
		if error != nil {
			log.Fatal(error)
		}
		rateResult := make(map[string]float64)
		for _,cubes := range Cubes {
			dateCubes.Date = cubes.Time
			for _,cube := range cubes.Cubes {
				s, _ := strconv.ParseFloat(cube.Rate, 64)
				rateResult[cube.Currency] = s
			
				dateCubes.Rates = rateResult
			}
			allcubes = append(allcubes,dateCubes)
		}
		c.JSON(http.StatusOK, allcubes)
	})
	router.Run(":8080")
}
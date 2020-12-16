package controllers

import (
	"context"
	"fmt"
	"log"
	"strconv"

	// "os"

	db "api/database"
	models "api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.mongodb.org/mongo-driver/bson"
)
func GetNewestRates(c *gin.Context){
	cubeCollection := db.Connector
	var Cubes []models.Cubes
	var dateCubes models.DateCubes
	data,_ := cubeCollection.Find(context.Background(),bson.M{})
	allcubes := make([]models.DateCubes,0)
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



	c.JSON(http.StatusOK, allcubes[0])
}

func GetRandomRates(c *gin.Context){
	cubeCollection := db.Connector
	var time models.Time
	if err := c.ShouldBindBodyWith(&time, binding.JSON); err != nil {
		log.Printf("%+v", err)
	}
	if time.Time != "" {
		var Cubes []models.Cubes
		var DateCubes models.DateCubes
		data,_ := cubeCollection.Find(context.Background(),bson.M{"time" : time.Time})
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
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pass body please"})
	}
	
} 

func GetPropertyOfAll(c *gin.Context){
	cubeCollection := db.Connector
	var Cubes []models.Cubes
	var dateCubes models.BigCubes
	
	allcubes := make([]models.BigCubes,0)

	data,_ := cubeCollection.Find(context.Background(),bson.M{})
	defer data.Close(context.Background())
	error := data.All(context.Background(),&Cubes)
	if error != nil {
		log.Fatal(error)
	}
	rateResult := make(map[string]float64)
	
	for _,cubes := range Cubes {
		for _,cube := range cubes.Cubes {
			s, _ := strconv.ParseFloat(cube.Rate, 64)
			rateResult[cube.Currency] = s
			dateCubes.Rates = rateResult
			}
			allcubes = append(allcubes, dateCubes)
		}
	// var max float64 = 0
	var valueCurrency models.ValuePerCurrency
	infoCurrency := make([]models.ValuePerCurrency,0)
	// var infoCurrency []models.ValuePerCurrency
	for i := 0; i < len(allcubes) -1; i++ {
		for k,v := range allcubes[i].Rates {
			fmt.Println(k, v)
			valueCurrency.Currency = k
			valueCurrency.MaxPerCurrency = v
			valueCurrency.MinPerCurrency = v
			valueCurrency.AveragePerCurrency = v
		}
		infoCurrency.append(infoCurrency,valueCurrency)
	}
	fmt.Println(infoCurrency)
	c.JSON(http.StatusOK, valueCurrency)
}
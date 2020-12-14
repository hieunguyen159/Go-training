package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cube struct {
	Currency  string `bson:"currency"`
	Rate string `bson:"rate"`
}
type Cubes struct{
	ID        primitive.ObjectID           `bson:"_id"`
	Time      string                       `bson:"time"`
	Cubes 	[]Cube 				    `bson:"Cube"`
}

type DateCubes struct {
	Date string
	Rates map[string]float64
}

type ValuePerCurrency struct {
	MinPerCurrency map[string]int
	MaxPerCurrency map[string]int
	AveragePerCurrency map[string]int
}
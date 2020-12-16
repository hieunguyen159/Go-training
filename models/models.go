package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Time struct {
	Time string
}
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
type BigCubes struct {
	Rates map[string]float64
}

type ValuePerCurrency struct {
	Currency string
	MinPerCurrency float64
	MaxPerCurrency float64
	AveragePerCurrency float64
}
type Info struct {
	Info []ValuePerCurrency
}
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cube struct {
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
}
type Rates struct {
	Id    primitive.ObjectID `bson:"_id,omitempty"`
	Time  time.Time          `json:"time"`
	Cubes []Cube
}
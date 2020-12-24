package models

type Cube struct {
	Currency string  `xml:"currency,attr" bson:"currency" json:"currency"`
	Rate     float64 `xml:"rate,attr" bson:"rate" json:"rate"`
}

type Cubes struct {
	Time  string `xml:"time,attr" bson:"time"`
	Cubes []Cube `xml:"Cube" bson:"Cube"`
}

type BigCube struct {
	BigCube []Cubes `xml:"Cube"`
}
type Envelope struct {
	Envelope BigCube `xml:"Cube"`
}
type DateCube struct {
	Date  string             `json:"time"`
	Rates map[string]float64 `json:"rates"`
}

type Time struct {
	Time string
}

type ValuePerCurrency struct {
	Currency           string  `bson:"_id" json:"currency"`
	MinPerCurrency     float64 `bson:"min" json:"minPerCurrency"`
	MaxPerCurrency     float64 `bson:"max" json:"maxPerCurrency"`
	AveragePerCurrency float64 `bson:"avg" json:"averagePerCurrency"`
}

type Form struct {
	Email    string   `json:"email"`
	Receiver []string `json:"receiver"`
}
type Email struct {
	Email  string `json:"email" bson:"email"`
	Status string `json:"status" bson:"status"`
}

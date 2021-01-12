package rates

type Time struct {
	Time string `json:"time"`
}

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

package rates

type RatesResponse struct {
	Date  string             `json:"time"`
	Rates map[string]float64 `json:"rates"`
}

type ValuePerCurrency struct {
	Currency           string  `bson:"_id" json:"currency"`
	MinPerCurrency     float64 `bson:"min" json:"minPerCurrency"`
	MaxPerCurrency     float64 `bson:"max" json:"maxPerCurrency"`
	AveragePerCurrency float64 `bson:"avg" json:"averagePerCurrency"`
}

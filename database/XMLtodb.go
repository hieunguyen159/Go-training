package db

import (
	models "api/models"
	"context"
	"encoding/xml"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetXMLfile() {
	resp, err := http.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	var envelope models.Envelope
	if err = xml.NewDecoder(resp.Body).Decode(&envelope); err != nil {
		log.Println(err)
	}
	var ui interface{}
	for _, t := range envelope.Envelope.BigCube {
		if t.Time == "2020-09-29" {
			ui = t
		}
	}
	Connector.InsertOne(context.Background(), ui)
}
func AddDataDaily() {
	resp, err := http.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	var envelope models.Envelope
	if err = xml.NewDecoder(resp.Body).Decode(&envelope); err != nil {
		log.Println(err)
	}
	cubeCollection := Connector
	var Cubes []models.Cubes
	data, _ := cubeCollection.Find(context.Background(), bson.M{})
	defer data.Close(context.Background())
	error := data.All(context.Background(), &Cubes)
	if error != nil {
		log.Fatal("Error")
	}
	var dataDaily interface{}
	for i, t := range envelope.Envelope.BigCube {
		if t.Time == Cubes[len(Cubes)-1].Time {
			dataDaily = envelope.Envelope.BigCube[i-1]
		}
	}
	cubeCollection.InsertOne(context.Background(), dataDaily)
}

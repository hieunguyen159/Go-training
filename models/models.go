package models

import (
	"html"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

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
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Email    string             `json:"email" bson:"email"`
	Status   string             `json:"status" bson:"status"`
	Reminded bool               `json:"reminded" bson:"reminded"`
}
type LoggedInUser struct {
	Email    string   `json:"email" bson:"email"`
	Password string   `json:"password" bson:"password"`
	Roles    []string `json:"roles" bson:"roles"`
}
type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	Roles    []string           `json:"roles" bson:"roles"`
	Active   bool               `json:"active" bson:"active"`
}
type Response struct {
	Token string `json:"token"`
}
type Status struct {
	Status bool `json:"status" bson:"status"`
}
type Role struct {
	Roles []string `json:"roles" bson:"roles"`
}

func NewUserBson(e User) User {
	return User{
		ID:       primitive.NewObjectID(),
		Email:    e.Email,
		Password: e.Password,
		Roles:    e.Roles,
		Active:   true,
	}
}
func NewEmailBSon(e string) Email {
	return Email{
		ID:       primitive.NewObjectID(),
		Email:    e,
		Status:   "waiting",
		Reminded: true,
	}

}

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}

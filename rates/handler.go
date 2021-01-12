package rates

import (
	db "api/database"
	"context"
	"encoding/xml"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// HandlerInterface interface
type HandlerInterface interface {
	Rates(c *gin.Context)
	SyncData()
}

// Handler struct
type Handler struct {
	Usecase UsecaseInterface
}

// NewHandler function
func NewHandler(us UsecaseInterface) HandlerInterface {
	return &Handler{Usecase: us}
}
func (h *Handler) Rates(c *gin.Context) {
	prefix := strings.TrimPrefix(c.Request.URL.Path, "/rates/")
	c.Header("Content-Type", "application/json")
	if prefix != "latest" && prefix != "analyze" {
		_, err := time.Parse("2006-01-02", prefix)
		if err != nil {
			log.Println("GetRatesDate: invalid format", err)
			c.JSON(400, "GetRatesDate: date invalid format")
			return
		}
	}
	resp, err := h.Usecase.GetRates(prefix)
	if err != nil {
		log.Println("GetRates: invalid format", err)
		c.JSON(500, "GetRates: error server")
		return
	}
	c.JSON(200, resp)

}
func (h *Handler) SyncData() {
	dao, _ := db.LoadConfig()
	resp, err := http.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	var envelope Envelope
	if err = xml.NewDecoder(resp.Body).Decode(&envelope); err != nil {
		log.Println(err)
	}
	var ui interface{}
	for _, t := range envelope.Envelope.BigCube {
		if t.Time == "2020-09-29" {
			ui = t
		}
	}
	_, err = dao.CubesCollection.InsertOne(context.Background(), ui)
	if err == nil {
		log.Printf("error")
	}
}

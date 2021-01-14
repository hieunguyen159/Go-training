package rates

import (
	db "api/database"
	"context"
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"strings"
	"time"
)

// HandlerInterface interface
type HandlerInterface interface {
	Rates(w http.ResponseWriter, r *http.Request)
	SyncData()
}

// Handler struct
type Handler struct {
	Usecase UsecaseInterface
}
type ErrMsg struct {
	Code    int
	Message string
}

// NewHandler function
func NewHandler(us UsecaseInterface) HandlerInterface {
	return &Handler{Usecase: us}
}
func (h *Handler) Rates(w http.ResponseWriter, r *http.Request) {
	prefix := strings.TrimPrefix(r.URL.Path, "/rates/")
	w.Header().Set("Content-Type", "application/json")
	if prefix != "latest" && prefix != "analyze" {
		_, err := time.Parse("2006-01-02", prefix)
		if err != nil {
			log.Println("GetRatesDate: invalid format", err)
			json.NewEncoder(w).Encode(&ErrMsg{
				Code:    400,
				Message: "GetRatesDate: date invalid format",
			})
			// c.JSON(400, "GetRatesDate: date invalid format")
			return
		}
	}
	resp, err := h.Usecase.GetRates(prefix)
	if err != nil {
		log.Println("GetRates: invalid format", err)
		json.NewEncoder(w).Encode(&ErrMsg{
			Code:    500,
			Message: "GetRates: error server",
		})
		// c.JSON(500, "GetRates: error server")
		return
	}
	json.NewEncoder(w).Encode(&resp)

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

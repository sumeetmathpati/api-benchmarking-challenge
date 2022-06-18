package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const (
	HOST  = "0.0.0.0"
	PORT  = "5000"
	INDEX = "events"
)

type CountResponse struct {
	EventCount int `json:"event_count"`
}

func main() {

	app := fiber.New()
	var r map[string]interface{}

	ES_HOST := os.Getenv("ES_HOST")
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://" + ES_HOST + ":9200",
			"http://" + ES_HOST + ":9201",
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	app.Get("/", func(c *fiber.Ctx) error {

		// New record to insert
		id := uuid.New()
		res, err := es.Index(
			INDEX,
			strings.NewReader(`{"title" : "events"}`),
			es.Index.WithDocumentID(id.String()),
			es.Index.WithRefresh("true"),
		)
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}

		// Get count of documents.
		res, err = es.Count(es.Count.WithIndex(INDEX))
		json.NewDecoder(res.Body).Decode(&r)

		return c.JSON(CountResponse{
			EventCount: int(r["count"].(float64)),
		})
	})

	log.Fatal(app.Listen(HOST + ":" + PORT))
}

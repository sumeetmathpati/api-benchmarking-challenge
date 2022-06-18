package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gofiber/fiber/v2"
)

const (
	HOST  = "0.0.0.0"
	PORT  = "5000"
	INDEX = "events"
)

func main() {
	var r map[string]interface{}
	app := fiber.New()
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	app.Get("/", func(c *fiber.Ctx) error {

		res, err := es.Index(
			INDEX,
			strings.NewReader(`{"title" : "events"}`),
			es.Index.WithDocumentID("1"),
			es.Index.WithRefresh("true"),
		)
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}

		res, err = es.Count(es.Count.WithIndex(INDEX))
		json.NewDecoder(res.Body).Decode(&r)

		return c.JSON(struct {
			EventCount int `json:"EventCount"`
		}{
			EventCount: int(r["count"].(float64)),
		})
	})

	log.Fatal(app.Listen(HOST + ":" + PORT))
}

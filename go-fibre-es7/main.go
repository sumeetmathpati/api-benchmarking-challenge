package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

const (
	HOST = "localhost"
	PORT = "5000"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(HOST + ":" + PORT))
}

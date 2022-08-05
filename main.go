package main

import (
	"log"
	"web-api/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
    app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
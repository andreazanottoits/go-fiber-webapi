package main

import (
	"log"
	"web-api/database"
	"web-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
    app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/api/user", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/user/:id", routes.GetUser)
	app.Delete("/api/user/:id", routes.DeleteUser)

	app.Post("/api/car", routes.CreateCar)
	app.Get("/api/cars", routes.GetCars)
	app.Delete("/api/users/:id", routes.DeleteCar)

	log.Fatal(app.Listen(":3000"))
}
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hekx11/goapi/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, asdasd!")
	})

	app.Listen(":3000")
}

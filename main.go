package main

import (
	"practice-commerce/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := config.DatabaseInit()

	app := fiber.New()

	app.Get("/testApi", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "hello",
		})
	})

	app.Listen(":3000")
}

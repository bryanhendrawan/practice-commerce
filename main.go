package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/testApi", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "hello",
		})
	})

	app.Listen(":3000")
}

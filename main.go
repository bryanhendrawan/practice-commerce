package main

import (
	"practice-commerce/config"
	"practice-commerce/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := config.DatabaseInit()

	app := fiber.New()

	app.Use(app)
	routes.Setup(app)

	app.Listen(":3000")
}

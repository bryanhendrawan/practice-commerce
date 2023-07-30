package main

import (
	"practice-commerce/config"
	"practice-commerce/controller"
	"practice-commerce/model"
	"practice-commerce/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// init
	db := config.DatabaseInit()
	models := config.ModelInit(
		model.NewProductModel(db),
		model.NewOrderModel(db),
	)
	controllers := config.ControllerInit(
		controller.NewProductController(models.ProductModel),
		controller.NewOrderController(models.OrderModel),
	)

	// fiber app
	app := fiber.New()

	// routes
	routes.Setup(app, controllers)

	// server
	app.Listen(":3000")
}

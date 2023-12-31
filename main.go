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
		model.NewCartModel(db),
		model.NewMerchantModel(db),
	)
	controllers := config.ControllerInit(
		controller.NewProductController(models.ProductModel),
		controller.NewOrderController(models.OrderModel, models.CartModel, models.ProductModel),
		controller.NewMerchantController(models.MerchantModel),
	)

	// fiber app
	app := fiber.New()

	// routes
	routes.Setup(app, controllers)

	// server
	app.Listen(":3000")
}

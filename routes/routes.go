package routes

import (
	"practice-commerce/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/merchants/:id/login", controller.Login)
	app.Get("/merchants/:id/logout", controller.Logout)

	app.Get("/products", controller.GetProducts)
	app.Get("/products/:id", controller.GetProductByID)
	app.Post("/product", controller.CreateProduct)
	app.Put("/product", controller.UpdateProduct)
	app.Delete("/product", controller.DeleteProduct)

	app.Get("/orders", controller.GetOrders)
	app.Get("/orders/:id", controller.GetOrderByID)
	app.Post("/order", controller.CreateOrder)
	app.Patch("/order", controller.UpdateOrder)
	app.Delete("/order", controller.DeleteOrder)
}

package routes

import (
	"practice-commerce/config"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, controller config.Controllers) {
	// app.Post("/merchants/:id/login", controller.Login)
	// app.Get("/merchants/:id/logout", controller.Logout)

	app.Get("/products", controller.ProductController.GetProducts)
	app.Get("/products/:id", controller.ProductController.GetProductDetail)
	app.Post("/product", controller.ProductController.CreateProduct)
	app.Patch("/product", controller.ProductController.UpdateProduct)
	app.Delete("/product", controller.ProductController.DeleteProduct)

	app.Get("/orders", controller.OrderController.GetOrders)
	app.Get("/orders/:id", controller.OrderController.GetOrderDetail)
	app.Get("/order/carts", controller.OrderController.GetCarts)
	app.Post("/order/cart", controller.OrderController.AddToCart)
	app.Post("/order/checkout", controller.OrderController.CreateOrder)
	app.Post("/order/cancel", controller.OrderController.CancelOrder)
}

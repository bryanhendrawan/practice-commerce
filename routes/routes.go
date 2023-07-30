package routes

import (
	"practice-commerce/config"
	"practice-commerce/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, controller config.Controllers) {
	app.Post("/merchant/login", controller.MerchantController.Login)
	app.Delete("/merchant/logout", controller.MerchantController.Logout)

	app.Get("/products", middleware.Authorize(), controller.ProductController.GetProducts)
	app.Get("/products/:id", middleware.Authorize(), controller.ProductController.GetProductDetail)
	app.Post("/product", middleware.Authorize(), controller.ProductController.CreateProduct)
	app.Patch("/product", middleware.Authorize(), controller.ProductController.UpdateProduct)
	app.Delete("/product", middleware.Authorize(), controller.ProductController.DeleteProduct)

	app.Get("/orders", middleware.Authorize(), controller.OrderController.GetOrders)
	app.Get("/orders/:id", middleware.Authorize(), controller.OrderController.GetOrderDetail)
	app.Get("/order/carts", middleware.Authorize(), controller.OrderController.GetCarts)
	app.Post("/order/cart", middleware.Authorize(), controller.OrderController.AddToCart)
	app.Post("/order/checkout", middleware.Authorize(), controller.OrderController.CreateOrder)
	app.Post("/order/cancel", middleware.Authorize(), controller.OrderController.CancelOrder)
}

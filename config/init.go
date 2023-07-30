package config

import (
	"practice-commerce/controller"
	"practice-commerce/model"
)

type models struct {
	ProductModel model.ProductModel
	OrderModel   model.OrderModel
}

func ModelInit(productModel model.ProductModel, orderModel model.OrderModel) models {
	return models{
		ProductModel: productModel,
		OrderModel:   orderModel,
	}
}

type Controllers struct {
	ProductController controller.ProductController
	OrderController   controller.OrderController
}

func ControllerInit(productController controller.ProductController, orderController controller.OrderController) Controllers {
	return Controllers{
		ProductController: productController,
		OrderController:   orderController,
	}
}

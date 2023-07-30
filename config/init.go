package config

import (
	"practice-commerce/controller"
	"practice-commerce/model"
)

type models struct {
	ProductModel  model.ProductModel
	OrderModel    model.OrderModel
	CartModel     model.CartModel
	MerchantModel model.MerchantModel
}

func ModelInit(productModel model.ProductModel, orderModel model.OrderModel, cartModel model.CartModel, merchantModel model.MerchantModel) models {
	return models{
		ProductModel:  productModel,
		OrderModel:    orderModel,
		CartModel:     cartModel,
		MerchantModel: merchantModel,
	}
}

type Controllers struct {
	ProductController  controller.ProductController
	OrderController    controller.OrderController
	MerchantController controller.MerchantController
}

func ControllerInit(productController controller.ProductController, orderController controller.OrderController, merchantController controller.MerchantController) Controllers {
	return Controllers{
		ProductController:  productController,
		OrderController:    orderController,
		MerchantController: merchantController,
	}
}

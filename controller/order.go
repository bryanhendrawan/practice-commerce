package controller

import (
	"practice-commerce/model"

	"github.com/gofiber/fiber/v2"
)

type OrderController interface {
	CreateOrder(c *fiber.Ctx) error
	UpdateOrder(c *fiber.Ctx) error
	DeleteOrder(c *fiber.Ctx) error
	GetOrders(c *fiber.Ctx) error
	GetOrderDetail(c *fiber.Ctx) error
}

type order struct {
	OrderModel model.OrderModel
}

func NewOrderController(orderModel model.OrderModel) OrderController {
	return &order{
		OrderModel: orderModel,
	}
}

func (o *order) CreateOrder(c *fiber.Ctx) error {
	return nil
}

func (o *order) UpdateOrder(c *fiber.Ctx) error {
	return nil
}

func (o *order) DeleteOrder(c *fiber.Ctx) error {
	return nil
}

func (o *order) GetOrders(c *fiber.Ctx) error {
	return nil
}

func (o *order) GetOrderDetail(c *fiber.Ctx) error {
	return nil
}

package controller

import (
	"fmt"
	"strconv"

	"practice-commerce/entity"
	"practice-commerce/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type OrderController interface {
	AddToCart(c *fiber.Ctx) error
	GetCarts(c *fiber.Ctx) error
	CreateOrder(c *fiber.Ctx) error
	CancelOrder(c *fiber.Ctx) error
	GetOrders(c *fiber.Ctx) error
	GetOrderDetail(c *fiber.Ctx) error
}

type order struct {
	OrderModel   model.OrderModel
	CartModel    model.CartModel
	ProductModel model.ProductModel
}

func NewOrderController(orderModel model.OrderModel, cartModel model.CartModel, productModel model.ProductModel) OrderController {
	return &order{
		OrderModel:   orderModel,
		CartModel:    cartModel,
		ProductModel: productModel,
	}
}

func (o *order) AddToCart(c *fiber.Ctx) error {
	requestBody := new(entity.AddToCartRequest)

	if err := c.BodyParser(requestBody); err != nil {
		return c.Status(400).JSON(entity.CommonResponse{
			Message: "Failed to parse request body",
			Error:   err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(requestBody); err != nil {
		return c.Status(400).JSON(entity.CommonResponse{
			Message: "Validate error",
			Error:   err.Error(),
		})
	}

	// checking if the product is available for add to cart
	product, err := o.ProductModel.GetProductByID(requestBody.ProductID)
	if err != nil {
		return c.Status(500).JSON(entity.CommonResponse{
			Message: "Failed to get product",
			Error:   err.Error(),
		})
	}

	if product.Stock < requestBody.Quantity {
		return c.Status(400).JSON(entity.CommonResponse{
			Message: fmt.Sprintf("Quantity more than stock. Product id %d", product.ID),
		})
	}

	// checking if the product need add to cart / update
	carts, err := o.CartModel.GetCarts(entity.GetCartParam{
		ProductID:  requestBody.ProductID,
		MerchantID: requestBody.MerchantID,
	})
	if err != nil {
		return c.Status(500).JSON(entity.CommonResponse{
			Message: "Failed to get carts",
			Error:   err.Error(),
		})
	}

	// add new product to cart
	if len(carts) == 0 {
		addCart := entity.Cart{
			ProductID:   requestBody.ProductID,
			ProductName: product.Name,
			Quantity:    requestBody.Quantity,
			MerchantID:  requestBody.MerchantID,
		}

		err := o.CartModel.AddCart(addCart)
		if err != nil {
			return c.Status(500).JSON(entity.CommonResponse{
				Message: "Failed add product to cart",
				Error:   err.Error(),
			})
		}

		return c.Status(201).JSON(entity.CommonResponse{
			Message: "Success add product to cart",
		})
	}

	// update existing product in cart
	updateCart := entity.Cart{
		ID:          carts[0].ID,
		ProductName: product.Name,
		Quantity:    requestBody.Quantity,
	}

	err = o.CartModel.UpdateCart(updateCart)
	if err != nil {
		return c.Status(500).JSON(entity.CommonResponse{
			Message: "Failed update product in cart",
			Error:   err.Error(),
		})
	}

	return c.Status(200).JSON(entity.CommonResponse{
		Message: "Success update product in cart",
	})
}

func (o *order) GetCarts(c *fiber.Ctx) error {
	param := new(entity.GetCartParam)

	if err := c.QueryParser(param); err != nil {
		return c.Status(400).JSON(entity.CommonResponse{
			Message: "Failed to parse cart param",
			Error:   err.Error(),
		})
	}

	if param.MerchantID < 1 {
		return c.Status(400).JSON(entity.CommonResponse{
			Message: "Merchant id can't empty",
		})
	}

	carts, err := o.CartModel.GetCarts(*param)
	if err != nil {
		return c.Status(500).JSON(entity.CommonResponse{
			Message: "Failed to get carts",
			Error:   err.Error(),
		})
	}

	return c.Status(200).JSON(entity.CartResponse{
		CommonResponse: entity.CommonResponse{
			Message: "Success get cart",
		},
		Data: carts,
	})
}

func (o *order) CreateOrder(c *fiber.Ctx) error {
	requestBody := new(entity.CreateOrderRequest)

	if err := c.BodyParser(requestBody); err != nil {
		return c.Status(400).JSON(entity.CommonResponse{
			Message: "Failed to parse request body",
			Error:   err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(requestBody); err != nil {
		return c.Status(400).JSON(entity.CommonResponse{
			Message: "Validate error",
			Error:   err.Error(),
		})
	}

	carts, err := o.CartModel.GetCarts(entity.GetCartParam{
		MerchantID: requestBody.MerchantID,
	})
	if err != nil {
		return c.Status(500).JSON(entity.CommonResponse{
			Message: "Failed to get carts",
			Error:   err.Error(),
		})
	}

	if len(carts) == 0 {
		return c.Status(400).JSON(entity.CommonResponse{
			Message: "Please add product to cart",
		})
	}

	createOrder := entity.Order{
		MerchantID: requestBody.MerchantID,
		Status:     entity.OrderCreated,
	}
	orderDetails := make([]entity.OrderDetail, 0)
	grandTotal := 0
	for _, cart := range carts {
		product, err := o.ProductModel.GetProductByID(cart.ProductID)
		if err != nil {
			return c.Status(500).JSON(entity.CommonResponse{
				Message: "Failed to get product detail",
				Error:   err.Error(),
			})
		}

		if product.Stock < cart.Quantity {
			return c.Status(400).JSON(entity.CommonResponse{
				Message: fmt.Sprintf("Quantity more than stock. Product id %d", product.ID),
			})
		}

		orderDetail := entity.OrderDetail{
			ProductID:     cart.ProductID,
			ProductName:   cart.ProductName,
			Quantity:      cart.Quantity,
			Price:         product.Price,
			TotalPrice:    cart.Quantity * product.Price,
			NewQtyProduct: product.Stock - cart.Quantity,
		}
		orderDetails = append(orderDetails, orderDetail)

		grandTotal += orderDetail.TotalPrice
	}
	createOrder.GrandTotal = grandTotal
	createOrder.OrderDetail = orderDetails

	order, err := o.OrderModel.CreateOrder(createOrder)
	if err != nil {
		return c.Status(500).JSON(entity.CommonResponse{
			Message: "Failed to get create order",
			Error:   err.Error(),
		})
	}

	return c.Status(200).JSON(entity.OrderResponse{
		CommonResponse: entity.CommonResponse{
			Message: "Success create order",
		},
		Data: &order,
	})
}

func (o *order) CancelOrder(c *fiber.Ctx) error {
	requestBody := new(entity.CancelOrderRequest)

	if err := c.BodyParser(requestBody); err != nil {
		return c.Status(400).JSON(entity.CommonResponse{
			Message: "Failed to parse request body",
			Error:   err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(requestBody); err != nil {
		return c.Status(400).JSON(entity.CommonResponse{
			Message: "Validate error",
			Error:   err.Error(),
		})
	}

	updateOrder := entity.Order{
		ID:     requestBody.OrderID,
		Status: entity.OrderCanceled,
	}

	err := o.OrderModel.UpdateOrder(updateOrder)
	if err != nil {
		return c.Status(500).JSON(entity.CommonResponse{
			Message: "Failed to cancel order",
			Error:   err.Error(),
		})
	}

	return c.Status(200).JSON(entity.CommonResponse{
		Message: "Success cancel order",
	})
}

func (o *order) GetOrders(c *fiber.Ctx) error {
	param := new(entity.GetOrderParam)

	if err := c.QueryParser(param); err != nil {
		return c.Status(400).JSON(entity.CommonResponse{
			Message: "Failed to parse order param",
			Error:   err.Error(),
		})
	}

	orders, err := o.OrderModel.GetOrders(*param)
	if err != nil {
		return c.Status(500).JSON(entity.CommonResponse{
			Message: "Failed to get orders",
			Error:   err.Error(),
		})
	}

	for i, order := range orders {
		orders[i].OrderDetail, err = o.OrderModel.GetOrderDetailByOrderID(order.ID)
		if err != nil {
			return c.Status(500).JSON(entity.CommonResponse{
				Message: "Failed to get order detail",
				Error:   err.Error(),
			})
		}
	}

	return c.Status(200).JSON(entity.OrderResponse{
		CommonResponse: entity.CommonResponse{
			Message: "Success get orders",
		},
		Datas: orders,
	})
}

func (o *order) GetOrderDetail(c *fiber.Ctx) error {
	requestID := c.Params("id")
	if requestID == "" {
		return c.Status(400).JSON(entity.CommonResponse{
			Message: "Order id can't empty",
		})
	}

	orderID, err := strconv.Atoi(requestID)
	if err != nil {
		return c.Status(400).JSON(entity.CommonResponse{
			Message: "Failed to parse order id",
			Error:   err.Error(),
		})
	}

	order, err := o.OrderModel.GetOrderByID(orderID)
	if err != nil {
		return c.Status(500).JSON(entity.CommonResponse{
			Message: "Failed to get order",
			Error:   err.Error(),
		})
	}

	order.OrderDetail, err = o.OrderModel.GetOrderDetailByOrderID(orderID)
	if err != nil {
		return c.Status(500).JSON(entity.CommonResponse{
			Message: "Failed to get order detail",
			Error:   err.Error(),
		})
	}

	return c.Status(200).JSON(entity.OrderResponse{
		CommonResponse: entity.CommonResponse{
			Message: "Success get order detail",
		},
		Data: &order,
	})
}

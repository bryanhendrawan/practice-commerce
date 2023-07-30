package controller

import (
	"strconv"

	"practice-commerce/entity"
	"practice-commerce/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	CreateProduct(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
	GetProducts(c *fiber.Ctx) error
	GetProductDetail(c *fiber.Ctx) error
}

type product struct {
	ProductModel model.ProductModel
}

func NewProductController(productModel model.ProductModel) ProductController {
	return &product{
		ProductModel: productModel,
	}
}

func (p *product) CreateProduct(c *fiber.Ctx) error {
	requestBody := new(entity.CreateProductRequest)

	if err := c.BodyParser(requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entity.CommonResponse{
			Message: "Failed to parse request body",
			Error:   err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entity.CommonResponse{
			Message: "Validate error",
			Error:   err.Error(),
		})
	}

	createProduct := entity.Product{
		Name:       requestBody.Name,
		Stock:      requestBody.Stock,
		Price:      requestBody.Price,
		MerchantID: requestBody.MerchantID,
	}

	product, err := p.ProductModel.CreateProduct(createProduct)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.CommonResponse{
			Message: "Failed to create product",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(entity.ProductResponse{
		CommonResponse: entity.CommonResponse{
			Message: "Success create product",
		},
		Data: &product,
	})
}

func (p *product) UpdateProduct(c *fiber.Ctx) error {
	requestBody := new(entity.UpdateProductRequest)

	if err := c.BodyParser(requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entity.CommonResponse{
			Message: "Failed to parse request body",
			Error:   err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entity.CommonResponse{
			Message: "Validate error",
			Error:   err.Error(),
		})
	}

	updateProduct := entity.Product{
		ID:    requestBody.ID,
		Name:  requestBody.Name,
		Stock: requestBody.Stock,
		Price: requestBody.Price,
	}

	err := p.ProductModel.UpdateProduct(updateProduct)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.CommonResponse{
			Message: "Failed to update product",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(entity.CommonResponse{
		Message: "Success update product",
	})
}

func (p *product) DeleteProduct(c *fiber.Ctx) error {
	requestBody := new(entity.DeleteProductRequest)

	if err := c.BodyParser(requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entity.CommonResponse{
			Message: "Failed to parse request body",
			Error:   err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entity.CommonResponse{
			Message: "Validate error",
			Error:   err.Error(),
		})
	}

	deleteProduct := entity.Product{
		ID: requestBody.ID,
	}

	if err := p.ProductModel.DeleteProduct(deleteProduct); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.CommonResponse{
			Message: "Failed to delete product",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(entity.CommonResponse{
		Message: "Success delete product",
	})
}

func (p *product) GetProducts(c *fiber.Ctx) error {
	param := new(entity.GetProductParam)

	if err := c.QueryParser(param); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entity.CommonResponse{
			Message: "Failed to parse product param",
			Error:   err.Error(),
		})
	}

	products, err := p.ProductModel.GetProducts(*param)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.CommonResponse{
			Message: "Failed to get products",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(entity.ProductResponse{
		CommonResponse: entity.CommonResponse{
			Message: "Success get products",
		},
		Datas: products,
	})
}

func (p *product) GetProductDetail(c *fiber.Ctx) error {
	requestID := c.Params("id")
	if requestID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(entity.CommonResponse{
			Message: "Product id can't empty",
		})
	}

	productID, err := strconv.Atoi(requestID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entity.CommonResponse{
			Message: "Failed to parse product id",
			Error:   err.Error(),
		})
	}

	product, err := p.ProductModel.GetProductByID(productID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.CommonResponse{
			Message: "Failed to get product detail",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(entity.ProductResponse{
		CommonResponse: entity.CommonResponse{
			Message: "Success get product detail",
		},
		Data: &product,
	})
}

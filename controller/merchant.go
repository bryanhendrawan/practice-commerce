package controller

import (
	"practice-commerce/entity"
	"practice-commerce/middleware"
	"practice-commerce/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type MerchantController interface {
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

type merchant struct {
	MerchantModel model.MerchantModel
}

func NewMerchantController(merchantModel model.MerchantModel) MerchantController {
	return &merchant{
		MerchantModel: merchantModel,
	}
}

func (m *merchant) Login(c *fiber.Ctx) error {
	requestBody := new(entity.LoginRequest)

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

	merchant, err := m.MerchantModel.GetMerchant(requestBody.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.CommonResponse{
			Message: "Failed to get merchant",
			Error:   err.Error(),
		})
	}

	if merchant.Password != requestBody.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(entity.CommonResponse{
			Message: "Password not match",
		})
	}

	tokenStr, err := middleware.CreateToken(merchant.ID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(entity.CommonResponse{
			Message: "Token is invalid or expired",
		})
	}

	merchant.Status = entity.StatusOnline
	err = m.MerchantModel.UpdateMerchant(merchant)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.CommonResponse{
			Message: "Failed to update merchant",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(entity.MerchantResponse{
		CommonResponse: entity.CommonResponse{
			Message: "Login success",
		},
		Token: tokenStr,
	})
}

func (m *merchant) Logout(c *fiber.Ctx) error {
	requestBody := new(entity.LogoutRequest)

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

	merchant := entity.Merchant{
		ID:     requestBody.ID,
		Status: entity.StatusOfflilne,
	}
	err := m.MerchantModel.UpdateMerchant(merchant)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entity.CommonResponse{
			Message: "Failed to update merchant",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(entity.CommonResponse{
		Message: "Logout success",
	})
}

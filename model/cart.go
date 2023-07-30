package model

import (
	"practice-commerce/entity"

	"gorm.io/gorm"
)

type CartModel interface {
	AddCart(cart entity.Cart) error
	UpdateCart(cart entity.Cart) error
	RemoveCart(cart entity.Cart) error
	GetCarts(param entity.GetCartParam) ([]entity.Cart, error)
}

type cart struct {
	db *gorm.DB
}

func NewCartModel(db *gorm.DB) CartModel {
	return &cart{db: db}
}

func (c *cart) AddCart(cart entity.Cart) error {
	result := c.db.Create(&cart)

	return result.Error
}

func (c *cart) UpdateCart(cart entity.Cart) error {
	result := c.db.Updates(&cart)

	return result.Error
}

func (c *cart) RemoveCart(cart entity.Cart) error {
	result := c.db.Delete(&cart)

	return result.Error
}

func (c *cart) GetCarts(param entity.GetCartParam) ([]entity.Cart, error) {
	var carts []entity.Cart
	queryDB := c.db

	if param.MerchantID > 0 {
		queryDB = queryDB.Where("merchant_id = ?", param.MerchantID)
	}
	if param.ProductID > 0 {
		queryDB = queryDB.Where("product_id = ?", param.ProductID)
	}

	result := queryDB.Order("created_at asc").Find(&carts)

	return carts, result.Error
}

package model

import (
	"practice-commerce/entity"

	"gorm.io/gorm"
)

type ProductModel interface {
	CreateProduct(product entity.Product) (entity.Product, error)
	UpdateProduct(product entity.Product) error
	DeleteProduct(product entity.Product) error
	GetProducts(param entity.GetProductParam) ([]entity.Product, error)
	GetProductByID(id int) (entity.Product, error)
}

type product struct {
	db *gorm.DB
}

func NewProductModel(db *gorm.DB) ProductModel {
	return &product{db: db}
}

func (p *product) CreateProduct(product entity.Product) (entity.Product, error) {
	result := p.db.Create(&product)

	return product, result.Error
}

func (p *product) UpdateProduct(product entity.Product) error {
	result := p.db.Updates(&product)

	return result.Error
}

func (p *product) DeleteProduct(product entity.Product) error {
	result := p.db.Delete(&product)

	return result.Error
}

func (p *product) GetProducts(param entity.GetProductParam) ([]entity.Product, error) {
	var products []entity.Product
	queryDB := p.db

	if param.MerchantID > 0 {
		queryDB = queryDB.Where("merchant_id = ?", param.MerchantID)
	}
	if param.Name != "" {
		queryDB = queryDB.Where("name LIKE ?", "%"+param.Name+"%")
	}

	result := queryDB.Order("updated_at desc").Find(&products)

	return products, result.Error
}

func (p *product) GetProductByID(id int) (entity.Product, error) {
	var product entity.Product
	result := p.db.Where("id = ?", id).First(&product)

	return product, result.Error
}

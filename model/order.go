package model

import (
	"practice-commerce/entity"

	"gorm.io/gorm"
)

type OrderModel interface {
	CreateOrder(order entity.Order) (entity.Order, error)
	UpdateOrder(order entity.Order) error
	DeleteOrderByID(id int) error
	GetOrders() ([]entity.Order, error)
	GetOrderByID(id int) (entity.Order, error)
}

type order struct {
	db *gorm.DB
}

func NewOrderModel(db *gorm.DB) OrderModel {
	return &order{db: db}
}

func (o *order) CreateOrder(order entity.Order) (entity.Order, error) {
	return entity.Order{}, nil
}

func (o *order) UpdateOrder(order entity.Order) error {
	return nil
}

func (o *order) DeleteOrderByID(id int) error {
	return nil
}

func (o *order) GetOrders() ([]entity.Order, error) {
	return []entity.Order{}, nil
}

func (o *order) GetOrderByID(id int) (entity.Order, error) {
	return entity.Order{}, nil
}

package model

import (
	"practice-commerce/entity"

	"gorm.io/gorm"
)

type OrderModel interface {
	CreateOrder(order entity.Order) (entity.Order, error)
	UpdateOrder(order entity.Order) error
	DeleteOrder(order entity.Order) error
	GetOrders(param entity.GetOrderParam) ([]entity.Order, error)
	GetOrderByID(id int) (entity.Order, error)
	GetOrderDetailByOrderID(orderID int) ([]entity.OrderDetail, error)
}

type order struct {
	db *gorm.DB
}

func NewOrderModel(db *gorm.DB) OrderModel {
	return &order{db: db}
}

func (o *order) CreateOrder(order entity.Order) (entity.Order, error) {
	trxErr := o.db.Transaction(func(tx *gorm.DB) error {
		for _, orderDetail := range order.OrderDetail {
			updateProduct := entity.Product{
				ID:    orderDetail.ProductID,
				Stock: orderDetail.NewQtyProduct,
			}
			if err := tx.Updates(&updateProduct).Error; err != nil {
				return err
			}
		}

		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		if err := tx.Where("merchant_id = ?", order.MerchantID).Delete(&entity.Cart{}).Error; err != nil {
			return err
		}

		return nil
	})

	return order, trxErr
}

func (o *order) UpdateOrder(order entity.Order) error {
	result := o.db.Updates(&order)

	return result.Error
}

func (o *order) DeleteOrder(order entity.Order) error {
	result := o.db.Delete(&order)

	return result.Error
}

func (o *order) GetOrders(param entity.GetOrderParam) ([]entity.Order, error) {
	var orders []entity.Order
	queryDB := o.db

	if param.MerchantID > 0 {
		queryDB = queryDB.Where("merchant_id = ?", param.MerchantID)
	}

	result := queryDB.Order("updated_at desc").Find(&orders)

	return orders, result.Error
}

func (o *order) GetOrderByID(id int) (entity.Order, error) {
	var order entity.Order
	result := o.db.Where("id = ?", id).First(&order)

	return order, result.Error
}

func (o *order) GetOrderDetailByOrderID(orderID int) ([]entity.OrderDetail, error) {
	var orderDetails []entity.OrderDetail
	result := o.db.Where("order_id = ?", orderID).Find(&orderDetails)

	return orderDetails, result.Error
}

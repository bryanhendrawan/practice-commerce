package entity

import "time"

type Order struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	MerchantID int       `json:"merchant_id"`
	GrandTotal int       `json:"grand_total"`
	Status     int       `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OrderDetail struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	OrderID     int       `json:"order_id"`
	ProductID   int       `json:"product_id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	Price       int       `json:"price"`
	TotalPrice  int       `json:"total_price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

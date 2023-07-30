package entity

import "time"

type Order struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	ProductID   int       `json:"product_id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	Price       int       `json:"price"`
	TotalPrice  int       `json:"total_price"`
	MerchantID  int       `json:"merchant_id"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

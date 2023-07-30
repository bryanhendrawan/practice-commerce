package entity

import "time"

const (
	OrderCreated   = 1
	OrderShipped   = 2
	OrderCompleted = 3
	OrderCanceled  = 4
)

type Order struct {
	ID          int           `json:"id" gorm:"primaryKey"`
	MerchantID  int           `json:"merchant_id"`
	GrandTotal  int           `json:"grand_total"`
	Status      int           `json:"status"`
	OrderDetail []OrderDetail `json:"order_detail"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
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

	// helper
	NewQtyProduct int `json:"-" gorm:"-"`
}

type GetOrderParam struct {
	MerchantID int `query:"merchant_id"`
}

type CreateOrderRequest struct {
	MerchantID int `json:"merchant_id" validate:"required,min=1"`
}

type CancelOrderRequest struct {
	OrderID int `json:"order_id" validate:"required,min=1"`
}

type OrderResponse struct {
	Data  *Order  `json:"data,omitempty"`
	Datas []Order `json:"datas,omitempty"`
	CommonResponse
}

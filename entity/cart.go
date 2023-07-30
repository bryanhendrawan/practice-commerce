package entity

import "time"

type Cart struct {
	ID          int       `json:"id"`
	MerchantID  int       `json:"merchant_id"`
	ProductID   int       `json:"product_id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetCartParam struct {
	ProductID  int `query:"product_id"`
	MerchantID int `query:"merchant_id"`
}

type AddToCartRequest struct {
	ProductID  int `json:"product_id" validate:"required,min=1"`
	Quantity   int `json:"quantity" validate:"required,min=1"`
	MerchantID int `json:"merchant_id" validate:"required,min=1"`
}

type CartResponse struct {
	Data []Cart `json:"data"`
	CommonResponse
}

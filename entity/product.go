package entity

import "time"

type Product struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	Stock      int       `json:"stock"`
	Price      int       `json:"price"`
	MerchantID int       `json:"merchant_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GetProductParam struct {
	Name       string `query:"name"`
	MerchantID int    `query:"merchant_id"`
}

type CreateProductRequest struct {
	Name       string `json:"name" validate:"required,min=1"`
	Stock      int    `json:"stock" validate:"required,min=0"`
	Price      int    `json:"price" validate:"required,min=1000"`
	MerchantID int    `json:"merchant_id" validate:"required,min=1"`
}

type UpdateProductRequest struct {
	ID    int    `json:"id" validate:"required,min=1"`
	Name  string `json:"name" validate:"required,min=1"`
	Stock int    `json:"stock" validate:"required,min=0"`
	Price int    `json:"price" validate:"required,min=1000"`
}

type DeleteProductRequest struct {
	ID int `json:"id" validate:"required,min=1"`
}

type ProductResponse struct {
	Data  *Product  `json:"data,omitempty"`
	Datas []Product `json:"datas,omitempty"`
	CommonResponse
}

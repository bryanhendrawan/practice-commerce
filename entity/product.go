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

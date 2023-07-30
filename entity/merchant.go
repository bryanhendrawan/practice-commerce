package entity

import "time"

const (
	StatusOnline   = "ONLINE"
	StatusOfflilne = "OFFLINE"
)

type Merchant struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=1"`
	Password string `json:"password" validate:"required,min=1"`
}

type LogoutRequest struct {
	ID int `json:"id" validate:"required,min=1"`
}

type MerchantResponse struct {
	Token string `json:"token"`
	CommonResponse
}

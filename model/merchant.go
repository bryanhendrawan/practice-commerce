package model

import (
	"practice-commerce/entity"

	"gorm.io/gorm"
)

type MerchantModel interface {
	UpdateMerchant(merchant entity.Merchant) error
	GetMerchant(username string) (entity.Merchant, error)
}

type merchant struct {
	db *gorm.DB
}

func NewMerchantModel(db *gorm.DB) MerchantModel {
	return &merchant{db: db}
}

func (m *merchant) UpdateMerchant(merchant entity.Merchant) error {
	result := m.db.Updates(&merchant)

	return result.Error
}

func (m *merchant) GetMerchant(username string) (entity.Merchant, error) {
	var merchant entity.Merchant

	result := m.db.Where(entity.Merchant{
		Username: username,
	}).First(&merchant)

	return merchant, result.Error
}

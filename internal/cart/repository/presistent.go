package repository

import (
	"pos/internal/cart"
	"pos/internal/model"

	"gorm.io/gorm"
)

type cartPresistentRepository struct {
	DB *gorm.DB
}

func NewcartPresistentRepository(db *gorm.DB) cart.CartRepository {
	return &cartPresistentRepository{
		DB: db,
	}
}

func (c cartPresistentRepository) Savetransaction(transaction *model.Cart) (err error) {
	err = c.DB.Create(transaction).Error
	return
}

func (c cartPresistentRepository) GetCartByTransactionID(transactionID uint) (carts []model.Cart, err error) {
	err = c.DB.Where("transaction_id = ?", transactionID).Find(&carts).Error
	return
}

package repository

import (
	"log"
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

func (c cartPresistentRepository) DeleteCart(cart uint) (err error) {
	err = c.DB.Where("id = ?", cart).Delete(&model.Cart{}).Error
	return
}

func (c cartPresistentRepository) UpdateCart(cart *model.Cart) (err error) {
	err = c.DB.Save(cart).Where("id=?", cart.ID).Error
	return
}

func (c cartPresistentRepository) SaveCart(cart *model.Cart) (err error) {
	log.Print(cart.ID)
	err = c.DB.Create(cart).Error
	return
}

func (c cartPresistentRepository) GetCartByTransactionID(offset, limit int64, transactionID uint) (carts []model.Cart, total int64, err error) {
	err = c.DB.Model(&model.Cart{}).Count(&total).Error
	if err != nil {
		return
	}
	err = c.DB.Where("transaction_id = ?", transactionID).Offset(int(offset)).Limit(int(limit)).Order("id desc").Find(&carts).Error
	if err != nil {
		return
	}
	return
}

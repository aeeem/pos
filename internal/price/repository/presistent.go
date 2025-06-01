package repository

import (
	"pos/internal/model"
	"pos/internal/price"

	"gorm.io/gorm"
)

type pricePresistentRepository struct {
	DB *gorm.DB
}

func NewPricePresistentRepository(db *gorm.DB) price.PriceRepository {

	return &pricePresistentRepository{
		DB: db,
	}
}

func (m pricePresistentRepository) SavePrice(price *model.Price) (err error) {
	err = m.DB.Create(price).Error
	return
}

func (m pricePresistentRepository) GetPrices(offset, limit int64, search string, itemID int64) (prices []model.Price, total int64, err error) {
	total = 0
	err = m.DB.Limit(int(limit)).Preload("Item").Offset(int(offset)).Where("item_id = ?", itemID).Find(&prices).Order("id desc").Count(&total).Error
	if err != nil {
		return
	}
	return
}

func (m pricePresistentRepository) GetPriceDetails(id int64) (price model.Price, err error) {
	err = m.DB.First(&price, id).Error
	return
}

func (m pricePresistentRepository) DeletePrice(id int64) (err error) {
	err = m.DB.Delete(&model.Price{}, id).Error
	return
}

func (m pricePresistentRepository) UpdatePrice(price *model.Price) (err error) {
	err = m.DB.Save(price).Error
	return
}

package repository

import (
	"pos/internal/domain"
	"pos/internal/price"

	"github.com/rs/zerolog/log"
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

func (m pricePresistentRepository) SavePrice(price *domain.Price) (err error) {
	err = m.DB.Create(price).Error
	return
}

func (m pricePresistentRepository) GetPrices(offset, limit int64, search string, itemID int64) (prices []domain.Price, total int64, err error) {
	total = 0
	err = m.DB.Limit(int(limit)).Preload("Item").Offset(int(offset)).Where("item_id = ?", itemID).Find(&prices).Order("id desc").Count(&total).Error
	if err != nil {
		return
	}
	return
}

func (m pricePresistentRepository) GetPriceDetails(id int64) (price domain.Price, err error) {
	err = m.DB.First(&price, id).Error
	return
}

func (m pricePresistentRepository) DeletePrice(id int64) (err error) {
	err = m.DB.Delete(&domain.Price{}, id).Error
	return
}

func (m pricePresistentRepository) UpdatePrice(price *domain.Price) (err error) {
	err = m.DB.Save(price).Where("id = ?", price.ID).Error
	log.Print(err)
	return
}

package repository

import (
	"pos/internal/domain"
	"pos/internal/item"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type itemPresistentRepository struct {
	DB *gorm.DB
}

func NewItemPresistenRepository(db *gorm.DB) item.ItemRepository {

	return &itemPresistentRepository{
		DB: db,
	}
}

func (m itemPresistentRepository) SaveItem(item *domain.Item) (err error) {
	err = m.DB.Create(item).Error
	return
}

func (m itemPresistentRepository) GetItems(offset, limit int64, search string) (items []domain.Item, total int64, err error) {
	err = m.DB.Model(&domain.Item{}).Count(&total).Error
	if err != nil {
		return
	}
	err = m.DB.Limit(int(limit)).Offset(int(offset)).Preload("Price").Order("id desc").Find(&items).Error
	if err != nil {
		return
	}
	log.Info(items)
	return
}

func (m itemPresistentRepository) GetItemDetails(id int64) (item domain.Item, err error) {
	err = m.DB.Preload("Price").First(&item, id).Error
	return
}
func (m itemPresistentRepository) DeleteItem(id int64) (err error) {
	err = m.DB.Delete(&domain.Item{}, id).Error
	return
}

func (m itemPresistentRepository) UpdateItem(item *domain.Item) (err error) {
	err = m.DB.Updates(item).Where("id = ?", item.ID).Model(&domain.Item{}).Error
	return
}

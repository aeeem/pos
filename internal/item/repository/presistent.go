package repository

import (
	"pos/internal/item"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type itemPresistentRepository struct {
	DB *gorm.DB
}

func NewItemPresistenRepository(db *gorm.DB) item.ItemRepository {
	if viper.GetString("seed") == "true" {
		SeedItem(db)
	}
	db.AutoMigrate(&item.Item{})
	return &itemPresistentRepository{
		DB: db,
	}
}

func (m itemPresistentRepository) SaveItem(item *item.Item) (err error) {
	err = m.DB.Create(item).Error
	return
}

func (m itemPresistentRepository) GetItems(offset, limit int64, search string) (items []item.Item, total int64, err error) {
	total = 0
	err = m.DB.Limit(int(limit)).Offset(int(offset)).Find(&items).Count(&total).Error
	if err != nil {
		return
	}
	return
}

func (m itemPresistentRepository) GetItemDetails(id int64) (item item.Item, err error) {
	return
}
func (m itemPresistentRepository) DeleteItem(id int64) (err error) {
	return
}

func (m itemPresistentRepository) UpdateItem(item *item.Item) (err error) {
	return
}

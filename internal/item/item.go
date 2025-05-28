package item

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemName     string `gorm:"index:idx_name,unique"`
	MaxPriceItem int64
}

type ItemUsecase interface {
	SaveItem(item *Item) (err error)
	GetItems(page, limit int64, search string) (items []Item, total int64, err error)
	GetItemDetails(id int64) (item Item, err error)
	DeleteItem(id int64) (err error)
	UpdateItem(item *Item) (err error)
}

type ItemRepository interface {
	SaveItem(item *Item) (err error)
	GetItems(offset, limit int64, search string) (items []Item, total int64, err error)
	GetItemDetails(id int64) (item Item, err error)
	DeleteItem(id int64) (err error)
	UpdateItem(item *Item) (err error)
}

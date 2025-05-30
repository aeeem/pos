package item

import "pos/internal/model"

type ItemUsecase interface {
	SaveItem(item *model.Item) (err error)
	GetItems(page, limit int64, search string) (items []model.Item, total int64, err error)
	GetItemDetails(id int64) (item model.Item, err error)
	DeleteItem(id int64) (err error)
	UpdateItem(item *model.Item) (err error)
}

type ItemRepository interface {
	SaveItem(item *model.Item) (err error)
	GetItems(offset, limit int64, search string) (items []model.Item, total int64, err error)
	GetItemDetails(id int64) (item model.Item, err error)
	DeleteItem(id int64) (err error)
	UpdateItem(item *model.Item) (err error)
}

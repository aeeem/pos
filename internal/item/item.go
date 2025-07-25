package item

import (
	"pos/internal/domain"
)

type ItemUsecase interface {
	SaveItem(item *domain.Item) (err error)
	GetItems(page, limit int64, search string) (items []domain.Item, total int64, err error)
	GetItemDetails(id int64) (item domain.Item, err error)
	DeleteItem(id int64) (err error)
	UpdateItem(item *domain.Item) (err error)
}

type ItemRepository interface {
	SaveItem(item *domain.Item) (err error)
	GetItems(offset, limit int64, search string) (items []domain.Item, total int64, err error)
	GetItemDetails(id int64) (item domain.Item, err error)
	DeleteItem(id int64) (err error)
	UpdateItem(item *domain.Item) (err error)
}

package usecase

import (
	"pos/internal/helper"
	"pos/internal/item"
	"pos/internal/model"
)

type itemUsecase struct {
	itemRepository item.ItemRepository
}

func NewItemUsecase(itemRepository item.ItemRepository) item.ItemUsecase {
	return &itemUsecase{
		itemRepository: itemRepository,
	}
}

func (m *itemUsecase) SaveItem(item *model.Item) (err error) {
	err = m.itemRepository.SaveItem(item)
	return
}

func (m *itemUsecase) GetItems(page, limit int64, search string) (items []model.Item, total int64, err error) {
	items, total, err = m.itemRepository.GetItems(helper.PageToOffset(page, limit), limit, search)
	return
}

func (m *itemUsecase) GetItemDetails(id int64) (item model.Item, err error) {
	item, err = m.itemRepository.GetItemDetails(id)
	return
}

func (m *itemUsecase) DeleteItem(id int64) (err error) {
	return
}

func (m *itemUsecase) UpdateItem(item *model.Item) (err error) {
	return
}

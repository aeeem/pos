package usecase

import (
	"pos/internal/domain"
	"pos/internal/helper"
	"pos/internal/item"
	"pos/internal/price"
)

type itemUsecase struct {
	itemRepository item.ItemRepository
	priceUsecase   price.PriceUsecase
}

func NewItemUsecase(itemRepository item.ItemRepository, priceUsecase price.PriceUsecase) item.ItemUsecase {
	return &itemUsecase{
		itemRepository: itemRepository,
		priceUsecase:   priceUsecase,
	}
}

func (m *itemUsecase) SaveItem(item *domain.Item) (err error) {
	prices := item.Price
	item.Price = []domain.Price{}
	err = m.itemRepository.SaveItem(item)
	if err != nil {
		return
	}
	for _, v := range prices {
		v.ItemID = item.ID
		v.Active = true
		err = m.priceUsecase.SavePrice(&v)
		if err != nil {
			return
		}
		item.Price = append(item.Price, v)
	}

	return
}

func (m *itemUsecase) GetItems(page, limit int64, search string) (items []domain.Item, total int64, err error) {
	items, total, err = m.itemRepository.GetItems(helper.PageToOffset(page, limit), limit, search)
	return
}

func (m *itemUsecase) GetItemDetails(id int64) (item domain.Item, err error) {
	item, err = m.itemRepository.GetItemDetails(id)
	return
}

func (m *itemUsecase) DeleteItem(id int64) (err error) {
	return
}

func (m *itemUsecase) UpdateItem(item *domain.Item) (err error) {
	price := item.Price

	item.Price = []domain.Price{}
	//get old image url
	old_item, err := m.GetItemDetails(int64(item.ID))
	if err != nil {
		return
	}

	if item.ImageUrl == "" {
		item.ImageUrl = old_item.ImageUrl
	}

	err = m.itemRepository.UpdateItem(item)
	if len(item.Price) > 0 {
		for _, v := range price {
			v.ItemID = item.ID
			v.Active = true
			err = m.priceUsecase.SavePrice(&v)
			if err != nil {
				return
			}
		}
	} else {
		item.Price = old_item.Price
	}
	*item, err = m.GetItemDetails(int64(item.ID))
	return
}

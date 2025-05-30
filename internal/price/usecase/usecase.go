package usecase

import (
	"pos/internal/helper"
	"pos/internal/item"
	"pos/internal/model"
	"pos/internal/price"
)

type priceUsecase struct {
	itemRepository  item.ItemRepository
	priceRepository price.PriceRepository
}

func NewPriceUsecase(priceRepository price.PriceRepository) price.PriceUsecase {
	return &priceUsecase{
		priceRepository: priceRepository,
	}
}

func (m *priceUsecase) SavePrice(price *model.Price) (err error) {
	err = m.priceRepository.SavePrice(price)
	return
}

// GetPrices gets a list of prices with pagination and search.
// It returns a slice of Price, total count of the search result, and error.
// The search parameter is used to search prices with their name or description.
func (m *priceUsecase) GetPrices(page, limit int64, search string) (prices []model.Price, total int64, err error) {
	prices, total, err = m.priceRepository.GetPrices(helper.PageToOffset(page, limit), limit, search)
	return
}
func (m *priceUsecase) GetPriceDetails(id int64) (price model.Price, err error) {
	return
}
func (m *priceUsecase) DeletePrice(id int64) (err error) {
	return
}
func (m *priceUsecase) UpdatePrice(price *model.Price) (err error) {
	return
}

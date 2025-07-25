package usecase

import (
	"pos/internal/domain"
	"pos/internal/helper"
	"pos/internal/item"
	"pos/internal/price"

	"github.com/rs/zerolog/log"
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

func (m *priceUsecase) SavePrice(price *domain.Price) (err error) {
	err = m.priceRepository.SavePrice(price)
	return
}

// GetPrices gets a list of prices with pagination and search.
// It returns a slice of Price, total count of the search result, and error.
// The search parameter is used to search prices with their name or description.
func (m *priceUsecase) GetPrices(page, limit int64, search string, itemID int64) (prices []domain.Price, total int64, err error) {
	prices, total, err = m.priceRepository.GetPrices(helper.PageToOffset(page, limit), limit, search, itemID)
	return
}
func (m *priceUsecase) GetPriceDetails(id int64) (price domain.Price, err error) {
	price, err = m.priceRepository.GetPriceDetails(id)
	return
}
func (m *priceUsecase) DeletePrice(id int64) (err error) {

	return m.priceRepository.DeletePrice(id)
}
func (m *priceUsecase) UpdatePrice(price *domain.Price) (err error) {
	log.Info().Any("price", price).Msg("price")
	err = m.priceRepository.UpdatePrice(price)
	return
}

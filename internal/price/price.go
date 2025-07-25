package price

import (
	"pos/internal/domain"
)

type PriceRepository interface {
	SavePrice(price *domain.Price) (err error)
	GetPrices(offset, limit int64, search string, itemID int64) (prices []domain.Price, total int64, err error)
	GetPriceDetails(id int64) (price domain.Price, err error)
	DeletePrice(id int64) (err error)
	UpdatePrice(price *domain.Price) (err error)
}

type PriceUsecase interface {
	SavePrice(price *domain.Price) (err error)
	GetPrices(page, limit int64, search string, itemID int64) (prices []domain.Price, total int64, err error)
	GetPriceDetails(id int64) (price domain.Price, err error)
	DeletePrice(id int64) (err error)
	UpdatePrice(price *domain.Price) (err error)
}

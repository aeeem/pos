package price

import (
	"pos/internal/model"
)

type PriceRepository interface {
	SavePrice(price *model.Price) (err error)
	GetPrices(offset, limit int64, search string, itemID int64) (prices []model.Price, total int64, err error)
	GetPriceDetails(id int64) (price model.Price, err error)
	DeletePrice(id int64) (err error)
	UpdatePrice(price *model.Price) (err error)
}

type PriceUsecase interface {
	SavePrice(price *model.Price) (err error)
	GetPrices(page, limit int64, search string, itemID int64) (prices []model.Price, total int64, err error)
	GetPriceDetails(id int64) (price model.Price, err error)
	DeletePrice(id int64) (err error)
	UpdatePrice(price *model.Price) (err error)
}

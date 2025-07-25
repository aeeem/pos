package delivery

import (
	"pos/internal/domain"
)

type GetItemResponse struct {
	Price []domain.Price `json:"price"`
	Total int64          `json:"total"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
}

type SavePriceResponse struct {
	Status string       `json:"status"`
	Price  domain.Price `json:"price"`
}

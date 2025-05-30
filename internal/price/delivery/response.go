package delivery

import "pos/internal/model"

type GetItemResponse struct {
	Price []model.Price `json:"price"`
	Total int64         `json:"total"`
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
}

type SavePriceResponse struct {
	Status string      `json:"status"`
	Price  model.Price `json:"price"`
}

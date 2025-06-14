package delivery

import "pos/internal/helper"

type SaveOrUpdate struct {
	Price  float64 `json:"price" validate:"required"`
	Unit   string  `json:"unit" validate:"required"`
	Active bool    `json:"active" validate:"required"`
	ItemID int64   `json:"item_id" validate:"required"`
}

type GetPriceRequest struct {
	helper.GetRequest
	ItemID int64 `json:"item_id" validate:"required"`
}

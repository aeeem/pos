package delivery

import "pos/internal/helper"

type SaveOrUpdate struct {
	Price  int64 `json:"price"`
	Active bool  `json:"active"`
	ItemID int64 `json:"item_id"`
}

type GetPriceRequest struct {
	helper.GetRequest
	ItemID int64 `json:"item_id" validate:"required"`
}

package delivery

import (
	"pos/internal/helper"
	"pos/internal/model"
)

type GetCartByTransactionIDResponse struct {
	Carts []model.Cart `json:"carts"`
	helper.ListResponse
}

type SaveCartResponse struct {
	Status string     `json:"status"`
	Cart   model.Cart `json:"cart"`
}

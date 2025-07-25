package delivery

import (
	"pos/internal/domain"
	"pos/internal/helper"
)

type GetCartByTransactionIDResponse struct {
	Carts []domain.Cart `json:"carts"`
	helper.ListResponse
}

type SaveCartResponse struct {
	Status string      `json:"status"`
	Cart   domain.Cart `json:"cart"`
}

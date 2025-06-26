package delivery

import (
	"pos/internal/cart/delivery"
	"pos/internal/helper"
)

type SaveOrUpdate struct {
	CustomerID   int                     `json:"customer_id" validate:"omitempty,numeric"`
	CustomerName string                  `json:"customer_name" validate:"required"`
	Status       *string                 `json:"status" validate:"oneof=completed cancelled pending draft"`
	Cart         []delivery.SaveOrUpdate `json:"cart" validate:"omitempty,dive"`
}

type GetTransactionsRequest struct {
	helper.GetRequest
	CustomerID int    `json:"customer_id" validate:"omitempty,numeric"`
	Status     string `json:"status,omitempty" validate:"omitempty,oneof=completed cancelled pending "`
}

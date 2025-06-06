package delivery

import "pos/internal/helper"

type SaveOrUpdate struct {
	CustomerID   int     `json:"customer_id" validate:"omitempty,numeric"`
	CustomerName string  `json:"customer_name" validate:"required"`
	Status       *string `json:"status" validate:"oneof=completed cancelled pending "`
}

type GetTransactionsRequest struct {
	helper.GetRequest
	CustomerID int    `json:"customer_id" validate:"omitempty,numeric"`
	Status     string `json:"status,omitempty" validate:"omitempty,oneof=completed cancelled pending "`
}

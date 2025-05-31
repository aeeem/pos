package delivery

import "pos/internal/helper"

type SaveOrUpdate struct {
	CustomerName string  `json:"customer_name" validate:"required"`
	Status       *string `json:"status" validate:"oneof=completed pending cancelled"`
}

type GetTransactionsRequest struct {
	helper.GetRequest
	Status *string `json:"status" validate:"omitempty,oneof=completed pending cancelleds"`
}

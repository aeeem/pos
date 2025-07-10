package delivery

import "pos/internal/helper"

type GetCustomeDebtRequest struct {
	helper.GetRequest
	CustomerID int    `json:"customer_id" validation:"required"`
	DateFrom   string `json:"date_from" validation:"omitempty,datetime"`
	DateTo     string `json:"date_to" validation:"omitempty,datetime"`
	DebtType   string `json:"debt_type" validation:"omitempty,oneof=paid unpaid"`
}
type Fetch struct {
	helper.GetRequest
	DateFrom string `json:"date_from" validation:"omitempty,datetime"`
	DateTo   string `json:"date_to" validation:"omitempty,datetime"`
	DebtType string `json:"debt_type" validation:"omitempty,oneof=paid unpaid"`
}

type PayCustomerDebtRequest struct {
	TransactionIDs []int   `json:"transaction_ids"`
	TotalAmount    float64 `json:"amount"`
}

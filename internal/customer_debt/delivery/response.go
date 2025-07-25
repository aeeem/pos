package delivery

import (
	"pos/internal/domain"
	"pos/internal/helper"
)

type GetResponse struct {
	helper.ListResponse
	Data []domain.CustomerDebt `json:"data"`
}

type PayDebtResponse struct {
	Transactions []domain.Transaction  `json:"transactions"`
	Debt         []domain.CustomerDebt `json:"debt"`
}

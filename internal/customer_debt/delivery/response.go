package delivery

import (
	"pos/internal/helper"
	"pos/internal/model"
)

type GetResponse struct {
	helper.ListResponse
	Data []model.CustomerDebt `json:"data"`
}

type PayDebtResponse struct {
	Transactions []model.Transaction  `json:"transactions"`
	Debt         []model.CustomerDebt `json:"debt"`
}

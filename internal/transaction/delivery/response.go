package delivery

import "pos/internal/model"

type GetItemResponse struct {
	Transaction []model.Transaction `json:"transactions"`
	Total       int64               `json:"total"`
	Page        int                 `json:"page"`
	Limit       int                 `json:"limit"`
}

type GetItemDetailResponse struct {
	Transaction model.Transaction `json:"transactions"`
	Total       int64             `json:"total"`
	Page        int               `json:"page"`
	Limit       int               `json:"limit"`
}

type SaveItemResponse struct {
	Status string            `json:"status"`
	Item   model.Transaction `json:"transaction"`
}

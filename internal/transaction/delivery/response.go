package delivery

import (
	"pos/internal/domain"
)

type GetItemResponse struct {
	Transaction []domain.Transaction `json:"transactions"`
	Total       int64                `json:"total"`
	Page        int                  `json:"page"`
	Limit       int                  `json:"limit"`
}

type GetItemDetailResponse struct {
	Transaction domain.Transaction `json:"transactions"`
	Total       int64              `json:"total"`
	Page        int                `json:"page"`
	Limit       int                `json:"limit"`
}

type SaveItemResponse struct {
	Status string             `json:"status"`
	Item   domain.Transaction `json:"transaction"`
}

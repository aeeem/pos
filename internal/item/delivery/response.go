package delivery

import (
	"pos/internal/domain"
)

type GetItemResponse struct {
	Items []domain.Item `json:"items"`
	Total int64         `json:"total"`
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
}

type GetItemDetailsResponse struct {
	Item domain.Item `json:"item"`
}

type SaveItemResponse struct {
	Status string      `json:"status"`
	Item   domain.Item `json:"item"`
}

type UpdateItemResponse struct {
	Status string      `json:"status"`
	Item   domain.Item `json:"item"`
}

package delivery

import "pos/internal/item"

type GetItemResponse struct {
	Items []item.Item `json:"items"`
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
}

type GetItemDetailsResponse struct {
	Item item.Item `json:"item"`
}

type SaveItemResponse struct {
	Status string    `json:"status"`
	Item   item.Item `json:"item"`
}

type UpdateItemResponse struct {
	Status string    `json:"status"`
	Item   item.Item `json:"item"`
}

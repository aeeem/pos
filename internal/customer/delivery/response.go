package delivery

import (
	"pos/internal/domain"
	"pos/internal/helper"
)

type CustomerListResponse struct {
	Customers []domain.Customer `json:"customers"`
	helper.ListResponse
}

type CustomerSaveOrUpdateResponse struct {
	Customer domain.Customer `json:"customer"`
	helper.StandardResponse
}

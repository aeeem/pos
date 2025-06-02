package delivery

import (
	"pos/internal/helper"
	"pos/internal/model"
)

type CustomerListResponse struct {
	Customers []model.Customer `json:"customers"`
	helper.ListResponse
}

type CustomerSaveOrUpdateResponse struct {
	Customer model.Customer `json:"customer"`
	helper.StandardResponse
}

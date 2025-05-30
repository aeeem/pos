package model

import "gorm.io/gorm"

type Price struct {
	gorm.Model
	Price  int64 `json:"price"`
	Active bool  `json:"active"`
	ItemID uint  `json:"item_id"`
	Item   Item  `json:"item"`
}

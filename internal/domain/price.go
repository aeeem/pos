package domain

import "gorm.io/gorm"

type Price struct {
	gorm.Model
	Unit   string  `json:"unit"`
	Price  float64 `json:"price"`
	Active bool    `json:"active"`
	ItemID uint    `json:"item_id"`
	Item   *Item   `json:"item" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}

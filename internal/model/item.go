package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemName     string  `gorm:"index:idx_name,unique" json:"item_name"`
	Price        []Price `gorm:"foreignKey:ItemID" json:"price"`
	MaxPriceItem int64   `json:"max_price_item"`
}

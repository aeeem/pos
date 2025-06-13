package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemName     string  `gorm:"index:idx_name,unique" json:"item_name" faker:"word,unique"`
	Price        []Price `gorm:"foreignKey:ItemID" json:"price" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	MaxPriceItem int64   `json:"max_price_item"`
}

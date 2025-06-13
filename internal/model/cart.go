package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ItemID        uint         `json:"item_id"`
	Item          *Item        `json:"item" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	TransactionID uint         `json:"transaction_id" `
	Transaction   *Transaction `json:"transaction" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	ItemName      string       `json:"item_name"`
	Unit          string       `json:"unit"`
	Quantity      int          `json:"quantity"`
	ItemPrice     int64        `json:"item_price"`
	PriceID       uint         `json:"price_id"`
	Price         Price        `json:"price" gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	SubPrice      int64        `json:"sub_price"`
}

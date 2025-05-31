package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ItemID        uint         `json:"item_id"`
	Item          *Item        `json:"item"`
	TransactionID uint         `json:"transaction_id"`
	Transaction   *Transaction `json:"transaction"`
	ItemName      string       `json:"item_name"`
	Quantity      int          `json:"quantity"`
	ItemPrice     int64        `json:"item_price"`
	PriceID       uint         `json:"price_id"`
	SubPrice      int64        `json:"sub_price"`
}

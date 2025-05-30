package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ItemID        uint
	Item          Item
	TransactionID uint
	Transaction   Transaction
	ItemName      string
	Quantity      int
	ItemPrice     int64
	PriceID       uint
	SubPrice      int64
}

package delivery

import "pos/internal/helper"

type GetCartByTransactionIDRequest struct {
	helper.GetRequest
	TransactionID int64 `json:"transaction_id" validate:"required"`
}

type SaveOrUpdate struct {
	ID            int64 `json:"id" validate:"-"`
	ItemID        uint  `json:"item_id"`
	TransactionID uint  `json:"transaction_id"`
	Quantity      int   `json:"quantity"`
	PriceID       uint  `json:"price_id"`
}

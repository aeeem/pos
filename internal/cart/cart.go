package cart

import "pos/internal/model"

type CartRepository interface {
	UpdateCart(transaction *model.Cart) (err error)
	GetCartByTransactionID(offset, limit int64, transactionID uint) (carts []model.Cart, total int64, err error)
	DeleteCart(transactionID uint) (err error)
	SaveCart(transaction *model.Cart) (err error)
}

type CartUsecase interface {
	UpdateCart(transaction *model.Cart) (err error)
	SaveCart(transaction *model.Cart) (err error)
	DeleteCart(transactionID uint) (err error)
	GetCartByTransactionID(page, limit int64, transactionID uint) (carts []model.Cart, total int64, err error) //[]model.Cart, error
}

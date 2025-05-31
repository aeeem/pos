package cart

import "pos/internal/model"

type CartRepository interface {
	GetCartByTransactionID(transactionID uint) (carts []model.Cart, err error) //[]model.Cart, error
	Savetransaction(transaction *model.Cart) (err error)
}

type CartUsecase interface {
	Savetransaction(transaction *model.Cart) (err error)
	GetCartByTransactionID(transactionID uint) (carts []model.Cart, err error) //[]model.Cart, error
}

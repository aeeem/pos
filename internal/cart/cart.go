package cart

import (
	"pos/internal/domain"
)

type CartRepository interface {
	UpdateCart(transaction *domain.Cart) (err error)
	GetCartByTransactionID(offset, limit int64, transactionID uint) (carts []domain.Cart, total int64, err error)
	DeleteCart(transactionID uint) (err error)
	SaveCart(transaction *domain.Cart) (err error)
}

type CartUsecase interface {
	UpdateCart(transaction *domain.Cart) (err error)
	SaveCart(transaction *domain.Cart) (err error)
	DeleteCart(transactionID uint) (err error)
	GetCartByTransactionID(page, limit int64, transactionID uint) (carts []domain.Cart, total int64, err error) //[]domain.Cart, error
}

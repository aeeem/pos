package usecase

import (
	"pos/internal/cart"
	"pos/internal/model"
)

type cartUsecase struct {
	CartRepository cart.CartRepository
}

func NewCartUsecase(cart cart.CartUsecase) cart.CartUsecase {
	return &cartUsecase{
		CartRepository: cart,
	}
}

func (c cartUsecase) Savetransaction(transaction *model.Cart) (err error) {
	err = c.CartRepository.Savetransaction(transaction)
	return
}
func (c cartUsecase) GetCartByTransactionID(transactionID uint) (carts []model.Cart, err error) {
	carts, err = c.CartRepository.GetCartByTransactionID(transactionID)

	return
}

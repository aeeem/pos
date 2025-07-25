package usecase

import (
	"pos/internal/cart"
	"pos/internal/domain"
	"pos/internal/helper"
	"pos/internal/item"
	"pos/internal/price"
)

type cartUsecase struct {
	CartRepository  cart.CartRepository
	ItemRepository  item.ItemRepository
	PriceRepository price.PriceRepository
}

func NewCartUsecase(
	cart cart.CartUsecase,
	ItemRepository item.ItemRepository,
	PriceRepository price.PriceRepository,
) cart.CartUsecase {
	return &cartUsecase{
		CartRepository:  cart,
		ItemRepository:  ItemRepository,
		PriceRepository: PriceRepository,
	}
}
func (c cartUsecase) DeleteCart(CartID uint) (err error) {
	err = c.CartRepository.DeleteCart(CartID)
	return
}
func (c cartUsecase) UpdateCart(Cart *domain.Cart) (err error) {
	err = c.CartRepository.SaveCart(Cart)
	return
}
func (c cartUsecase) SaveCart(cart *domain.Cart) (err error) {
	// get price
	items, err := c.ItemRepository.GetItemDetails(int64(cart.ItemID))
	if err != nil {
		return
	}

	prices, err := c.PriceRepository.GetPriceDetails(int64(cart.PriceID))
	if err != nil {
		return
	}

	cart.ItemName = items.ItemName
	cart.SubPrice = prices.Price * cart.Quantity
	cart.Unit = prices.Unit
	cart.ItemPrice = prices.Price
	err = c.CartRepository.SaveCart(cart)
	return
}
func (c cartUsecase) GetCartByTransactionID(page, limit int64, transactionID uint) (carts []domain.Cart, total int64, err error) {
	carts, total, err = c.CartRepository.GetCartByTransactionID(helper.PageToOffset(page, limit), limit, transactionID)

	return
}

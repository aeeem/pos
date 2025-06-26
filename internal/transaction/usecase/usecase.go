package usecase

import (
	"pos/internal/cart"
	"pos/internal/item"
	"pos/internal/price"

	"pos/internal/helper"
	"pos/internal/model"
	"pos/internal/transaction"
)

type transactionUsecase struct {
	ItemUsecase           item.ItemUsecase
	PriceUsecase          price.PriceUsecase
	CartUsecase           cart.CartUsecase
	transactionRepository transaction.TransactionRepository
}

func NewTransactionUsecase(tRepo transaction.TransactionRepository,
	cart cart.CartUsecase,
	ItemUsecase item.ItemUsecase,
	PriceUsecase price.PriceUsecase,
) transaction.TransactionUsecase {
	return &transactionUsecase{
		ItemUsecase:           ItemUsecase,
		PriceUsecase:          PriceUsecase,
		CartUsecase:           cart,
		transactionRepository: tRepo,
	}
}

func (t *transactionUsecase) Savetransaction(transaction *model.Transaction) (err error) {
	carts := transaction.Cart
	transaction.Cart = nil
	err = t.transactionRepository.Savetransaction(transaction)
	if err != nil {
		return
	}

	for _, v := range carts {
		//get item name

		v.TransactionID = transaction.ID
		err = t.CartUsecase.SaveCart(&v)
		if err != nil {
			return err
		}
	}
	transaction.Cart = make([]model.Cart, len(carts))
	transaction.Cart = carts
	return
}
func (t *transactionUsecase) GetTransactions(page, limit int64, search string, status model.Status, customerID int64) (transactions []model.Transaction, total int64, err error) {
	transactions, total, err = t.transactionRepository.GetTransactions(helper.PageToOffset(page, limit), limit, search, status, customerID)
	return
}
func (t *transactionUsecase) GetTransactionDetails(id int64) (transaction model.Transaction, err error) {
	transaction, err = t.transactionRepository.GetTransactionDetails(id)
	return
}
func (t *transactionUsecase) DeleteTransaction(id int64) (err error) {

	//check if transaction exist
	if _, err = t.GetTransactionDetails(id); err != nil {
		return
	}
	err = t.transactionRepository.Deletetransaction(id)
	return
}
func (t *transactionUsecase) UpdateTransaction(transaction *model.Transaction) (err error) {
	//check if transaction is exist
	oldtx, err := t.GetTransactionDetails(int64(transaction.ID))
	if err != nil {
		return
	}
	if transaction.CustomerTransactionNo != 0 {
		oldtx.CustomerTransactionNo = transaction.CustomerTransactionNo
	}

	if oldtx.CustomerID != transaction.CustomerID {
		oldtx.CustomerID = transaction.CustomerID
	}

	if oldtx.CustomerName != transaction.CustomerName {
		oldtx.CustomerName = transaction.CustomerName
	}

	if oldtx.Status != transaction.Status {
		oldtx.Status = transaction.Status
	}

	err = t.transactionRepository.Updatetransaction(&oldtx)
	return
}

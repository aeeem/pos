package usecase

import (
	"pos/internal/model"
	"pos/internal/transaction"

	"gorm.io/gorm"
)

type transactionUsecase struct {
	transactionRepository transaction.TransactionRepository
}

func NewTransactionUsecase(tRepo transaction.TransactionRepository) transaction.TransactionUsecase {
	return &transactionUsecase{
		transactionRepository: tRepo,
	}
}

func (t *transactionUsecase) Savetransaction(transaction *model.Transaction) (err error) {
	err = t.transactionRepository.Savetransaction(transaction)

	return
}
func (t *transactionUsecase) GetTransactions(page, limit int64, search string, status model.Status) (transactions []model.Transaction, total int64, err error) {
	transactions, total, err = t.transactionRepository.GetTransactions(page, limit, search, status)
	return
}
func (t *transactionUsecase) GetTransactionDetails(id int64) (transaction model.Transaction, err error) {
	transaction, err = t.transactionRepository.GetTransactionDetails(id)
	return
}
func (t *transactionUsecase) DeleteTransaction(id int64) (err error) {
	err = t.transactionRepository.Deletetransaction(id)
	return
}
func (t *transactionUsecase) UpdateTransaction(transaction *model.Transaction) (err error) {
	//check if transaction is exist
	_, err = t.GetTransactionDetails(int64(transaction.ID))
	if err != nil || err.Error() == gorm.ErrRecordNotFound.Error() {
		return
	}

	err = t.transactionRepository.Updatetransaction(transaction)
	return
}

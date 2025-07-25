package transaction

import (
	"pos/internal/domain"
)

type TransactionRepository interface {
	Savetransaction(transaction *domain.Transaction) (err error)

	GetTransactions(page, limit int64, search string, status domain.Status, customerID int64) (transactions []domain.Transaction, total int64, err error)
	GetTransactionDetails(id int64) (transaction domain.Transaction, err error)
	Deletetransaction(id int64) (err error)
	Updatetransaction(transaction *domain.Transaction) (err error)
}

type TransactionUsecase interface {
	Savetransaction(transaction *domain.Transaction) (err error)
	GetTransactions(page, limit int64, search string, status domain.Status, customerID int64) (transactions []domain.Transaction, total int64, err error)
	GetTransactionDetails(id int64) (transaction domain.Transaction, err error)
	DeleteTransaction(id int64) (err error)
	UpdateTransaction(transaction *domain.Transaction) (err error)
}

package transaction

import "pos/internal/model"

type TransactionRepository interface {
	Savetransaction(transaction *model.Transaction) (err error)
	GetTransactions(page, limit int64, search string, status model.Status) (transactions []model.Transaction, total int64, err error)
	GetTransactionDetails(id int64) (transaction model.Transaction, err error)
	Deletetransaction(id int64) (err error)
	Updatetransaction(transaction *model.Transaction) (err error)
}

type TransactionUsecase interface {
	Savetransaction(transaction *model.Transaction) (err error)
	GetTransactions(page, limit int64, search string, status model.Status) (transactions []model.Transaction, total int64, err error)
	GetTransactionDetails(id int64) (transaction model.Transaction, err error)
	DeleteTransaction(id int64) (err error)
	UpdateTransaction(transaction *model.Transaction) (err error)
}

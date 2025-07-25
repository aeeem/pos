package customerdebt

import (
	"pos/internal/domain"
)

type CustomerDebtRepository interface {
	GetDebtDetails(debtID uint) (res domain.CustomerDebt, err error)
	GetByCustomerID(customerID uint, debtType string, dateFrom string, dateTo string, page, limit int, search string) (res []domain.CustomerDebt, total int64, err error)
	Fetch(debtType string, dateFrom string, dateTo string, page, limit int, search string) (res []domain.CustomerDebt, total int64, err error)
	PayCustomerDebt(customerID int, status string, transaction_id int, amount float64) (err error)
	GetByTransactionID(transaction_id uint) (customerDebt domain.CustomerDebt, err error)
}

type CustomerDebtUsecase interface {
	GetDebtDetails(debtID uint) (res domain.CustomerDebt, err error)
	GetByCustomerID(customerID uint, debtType string, dateFrom string, dateTo string, page, limit int, search string) (res []domain.CustomerDebt, total int64, err error)
	Fetch(debtType string, dateFrom string, dateTo string, page, limit int, search string) (res []domain.CustomerDebt, total int64, err error)
	PayCustomerDebt(customerID int, transaction_id []int, amount float64) (Transactions []domain.Transaction, Debts []domain.CustomerDebt, err error)
}

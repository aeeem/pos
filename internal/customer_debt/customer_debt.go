package customerdebt

import "pos/internal/model"

type CustomerDebtRepository interface {
	GetDebtDetails(debtID uint) (res model.CustomerDebt, err error)
	GetByCustomerID(customerID uint, debtType string, dateFrom string, dateTo string, page, limit int, search string) (res []model.CustomerDebt, total int64, err error)
	Fetch(debtType string, dateFrom string, dateTo string, page, limit int, search string) (res []model.CustomerDebt, total int64, err error)
	PayCustomerDebt(customerID int, status string, transaction_id int, amount float64) (err error)
	GetByTransactionID(transaction_id uint) (customerDebt model.CustomerDebt, err error)
}

type CustomerDebtUsecase interface {
	GetDebtDetails(debtID uint) (res model.CustomerDebt, err error)
	GetByCustomerID(customerID uint, debtType string, dateFrom string, dateTo string, page, limit int, search string) (res []model.CustomerDebt, total int64, err error)
	Fetch(debtType string, dateFrom string, dateTo string, page, limit int, search string) (res []model.CustomerDebt, total int64, err error)
	PayCustomerDebt(customerID int, transaction_id []int, amount float64) (Transactions []model.Transaction, Debts []model.CustomerDebt, err error)
}

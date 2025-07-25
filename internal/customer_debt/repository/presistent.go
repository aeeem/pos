package repository

import (
	customerdebt "pos/internal/customer_debt"
	"pos/internal/domain"

	"gorm.io/gorm"
)

type CustomerDebtPresistentRepository struct {
	DB *gorm.DB
}

func NewCustomerDebtPresistentRepository(db *gorm.DB) customerdebt.CustomerDebtRepository {
	return &CustomerDebtPresistentRepository{
		DB: db,
	}
}

func (C *CustomerDebtPresistentRepository) GetDebtDetails(debtID uint) (res domain.CustomerDebt, err error) {
	queries := C.DB.Model(&domain.CustomerDebt{}).Preload("Mutation").Where("id = ?", debtID)
	err = queries.First(&res).Error
	return
}

func (C *CustomerDebtPresistentRepository) GetByCustomerID(customerID uint, debtType string, dateFrom string, dateTo string, page, limit int, search string) (res []domain.CustomerDebt, total int64, err error) {
	queries := C.DB.Model(&domain.CustomerDebt{}).Where("customer_id = ?", customerID)
	if debtType != "" {
		queries.Where("debt_status", debtType)
	}
	if dateFrom != "" {
		queries.Where("created_at >= ?", dateFrom)
	}
	if dateTo != "" {
		queries.Where("created_at <= ?", dateTo)
	}
	total = 0
	err = queries.Count(&total).Error
	if err != nil {
		return
	}
	err = queries.Preload("Customer").Limit(limit).Offset(page).Find(&res).Order("created_at desc").Error
	return
}

func (C *CustomerDebtPresistentRepository) GetByTransactionID(transaction_id uint) (customerDebt domain.CustomerDebt, err error) {
	queries := C.DB.Model(&domain.CustomerDebt{}).Where("trx_id = ?", transaction_id)
	err = queries.First(&customerDebt).Error
	return
}

func (C *CustomerDebtPresistentRepository) Fetch(debtType string, dateFrom string, dateTo string, page, limit int, search string) (res []domain.CustomerDebt, total int64, err error) {
	queries := C.DB.Model(&domain.CustomerDebt{})
	if debtType != "" {
		queries.Where("debt_status", debtType)
	}
	if dateFrom != "" {
		queries.Where("created_at >= ?", dateFrom)
	}
	if dateTo != "" {
		queries.Where("created_at <= ?", dateTo)
	}
	total = 0
	err = queries.Count(&total).Error
	if err != nil {
		return
	}
	err = queries.Preload("Customer").Offset(page).Limit(limit).Find(&res).Order("created_at desc").Error
	return
}
func (C *CustomerDebtPresistentRepository) PayCustomerDebt(customerID int, status string, transaction_id int, amount float64) (err error) {
	CustomerDebts := domain.CustomerDebt{
		CustomerID: uint(customerID),
		TrxID:      uint(transaction_id),
	}

	err = C.DB.Model(&CustomerDebts).
		Where("customer_id = ?", customerID).
		Where("debt_status = ? or debt_status = ?", "unpaid", "half_paid").
		Where("trx_id = ?", transaction_id).
		Updates(&map[string]interface{}{
			"debt_status":   status,
			"paid_amount":   gorm.Expr("paid_amount + ?", amount),
			"unpaid_amount": gorm.Expr("unpaid_amount - ?", amount),
		}).Error
	if err != nil {
		return
	}
	return
}

package repository

import (
	"pos/internal/domain"

	"gorm.io/gorm"

	"pos/internal/transaction"
)

type transactionRepository struct {
	DB *gorm.DB
}

func NewTransactionPresistentRepository(db *gorm.DB) transaction.TransactionRepository {
	db.AutoMigrate(&domain.Transaction{})
	return &transactionRepository{
		DB: db,
	}
}

func (t *transactionRepository) Savetransaction(transaction *domain.Transaction) (err error) {
	err = t.DB.Create(&transaction).Error
	if err != nil {
		return
	}
	return
}
func (t *transactionRepository) GetTransactions(page, limit int64, search string, status domain.Status, customerID int64) (transactions []domain.Transaction, total int64, err error) {
	total = int64(0)
	err = t.DB.Model(&domain.Transaction{}).Count(&total).Error
	err = t.DB.Limit(int(limit)).Offset(int(page)).Preload("Cart",
		func(db *gorm.DB) *gorm.DB {
			return db.Order("carts.id ASC")
		}).Where("customer_id = ?", customerID).Order("customer_transaction_no desc").Find(&transactions).Error
	if err != nil {
		return
	}
	return
}

func (t *transactionRepository) GetTransactionDetails(id int64) (transaction domain.Transaction, err error) {
	err = t.DB.Preload("Cart",
		func(db *gorm.DB) *gorm.DB {
			return db.Order("carts.id desc")
		}).First(&transaction, id).Error
	return
}
func (t *transactionRepository) Deletetransaction(id int64) (err error) {
	err = t.DB.Delete(&domain.Transaction{}, id).Error
	return
}
func (t *transactionRepository) Updatetransaction(transaction *domain.Transaction) (err error) {
	err = t.DB.Save(&transaction).Error
	return
}

package repository

import (
	"gorm.io/gorm"

	"pos/internal/model"
	"pos/internal/transaction"
)

type transactionRepository struct {
	DB *gorm.DB
}

func NewTransactionPresistentRepository(db *gorm.DB) transaction.TransactionRepository {
	db.AutoMigrate(&model.Transaction{})
	return &transactionRepository{
		DB: db,
	}
}

func (t *transactionRepository) Savetransaction(transaction *model.Transaction) (err error) {
	err = t.DB.Create(transaction).Error
	if err != nil {
		return
	}
	return
}
func (t *transactionRepository) GetTransactions(page, limit int64, search string, status model.Status) (transactions []model.Transaction, total int64, err error) {
	total = 0
	err = t.DB.Limit(int(limit)).Offset(int(page)).Preload("Cart").Find(&transactions).Count(&total).Error
	if err != nil {
		return
	}
	return
}

func (t *transactionRepository) GetTransactionDetails(id int64) (transaction model.Transaction, err error) {
	err = t.DB.Preload("Cart").First(&transaction, id).Error
	return
}
func (t *transactionRepository) Deletetransaction(id int64) (err error) {
	return
}
func (t *transactionRepository) Updatetransaction(transaction *model.Transaction) (err error) {
	return
}

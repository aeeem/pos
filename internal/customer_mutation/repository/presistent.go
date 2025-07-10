package repository

import (
	customermutation "pos/internal/customer_mutation"
	"pos/internal/model"

	"gorm.io/gorm"
)

type CustomerMutationRepository struct {
	DB *gorm.DB
}

func NewCustomerMutationRepository(db *gorm.DB) customermutation.CustomerMutationRepository {
	return &CustomerMutationRepository{
		DB: db,
	}
}
func (Cm *CustomerMutationRepository) SaveCustomerMutation(DebtMutation *model.CustomerDebtMutations) (err error) {
	err = Cm.DB.Model(&model.CustomerDebtMutations{}).Create(&DebtMutation).Error
	return
}

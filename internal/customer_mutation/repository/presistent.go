package repository

import (
	customermutation "pos/internal/customer_mutation"
	"pos/internal/domain"

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
func (Cm *CustomerMutationRepository) SaveCustomerMutation(DebtMutation *domain.CustomerDebtMutations) (err error) {
	err = Cm.DB.Model(&domain.CustomerDebtMutations{}).Create(&DebtMutation).Error
	return
}

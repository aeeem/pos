package repository

import (
	"pos/internal/model"
	"pos/internal/mutation"

	"gorm.io/gorm"
)

type MutationRepository struct {
	DB *gorm.DB
}

func NewMutationRepository(db *gorm.DB) mutation.MutationRepository {
	return &MutationRepository{
		DB: db,
	}

}

func (M *MutationRepository) GetCustomerMutation(customerID uint, MutationType string) (CustomerMutation []model.Mutation, err error) {
	err = M.DB.Where("customer_id = ?", customerID).Where("mutation_type = ?", MutationType).Find(&CustomerMutation).Error
	return
}

func (M *MutationRepository) SaveCustomerMutation(mutation *model.Mutation) (err error) {
	err = M.DB.Model(&model.Mutation{}).Create(&mutation).Error
	if err != nil {
		return
	}

	//updating join tables
	return
}

func (M *MutationRepository) UpdateCustomerMutation(mutationID uint, mutation *model.Mutation) (err error) {
	err = M.DB.Updates(&mutation).Where("id = ?", mutationID).Model(&model.Mutation{}).Error
	return
}

func (M *MutationRepository) GetCustomerBalance(customerID uint) (CustomerBalance float64, err error) {
	err = M.DB.Model(&model.Mutation{}).Where("customer_id = ?", customerID).Order("id DESC").Select("customer_balance").First(&CustomerBalance).Error
	return
}

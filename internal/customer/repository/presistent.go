package repository

import (
	"pos/internal/customer"
	"pos/internal/domain"

	"gorm.io/gorm"
)

type customerPresistentRepository struct {
	db *gorm.DB
}

func NewCustomerPresistentRepository(db *gorm.DB) customer.CustomerRepository {

	return &customerPresistentRepository{
		db: db,
	}

}

func (r customerPresistentRepository) SaveCustomer(customer *domain.Customer) (err error) {
	err = r.db.Create(&customer).Error
	return
}
func (r customerPresistentRepository) GetAllCustomer(offset, limit int64, search string) (customers []domain.Customer, total int64, err error) {
	err = r.db.Model(&domain.Customer{}).Count(&total).Error
	if err != nil {
		return
	}
	err = r.db.Limit(int(limit)).Offset(int(offset)).Where("customer_name LIKE ?", "%"+search+"%").Order("created_at desc").Find(&customers).Error
	if err != nil {
		return
	}
	return
}
func (r customerPresistentRepository) GetCustomerDetails(id int64) (customer domain.Customer, err error) {
	return
}
func (r customerPresistentRepository) DeleteCustomer(id int64) (err error) {
	return
}
func (r customerPresistentRepository) UpdateCustomer(customer *domain.Customer) (err error) {
	return
}

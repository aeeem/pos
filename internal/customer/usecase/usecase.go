package usecase

import (
	"pos/internal/customer"
	"pos/internal/helper"
	"pos/internal/model"
)

type customerUsecase struct {
	CustomerRepository customer.CustomerRepository
}

func NewCustomerUsecase(repository customer.CustomerRepository) customer.CustomerUsecase {
	return &customerUsecase{
		CustomerRepository: repository,
	}
}
func (m *customerUsecase) SaveCustomer(customer *model.Customer) (err error) {
	err = m.CustomerRepository.SaveCustomer(customer)
	return
}

func (m *customerUsecase) GetAllCustomer(page, limit int64, search string) (customers []model.Customer, total int64, err error) {
	customers, total, err = m.CustomerRepository.GetAllCustomer(helper.PageToOffset(page, limit), limit, search)
	return
}

//todo:implement
func (m *customerUsecase) GetCustomerDetails(id int64) (customer model.Customer, err error) {
	return
}
func (m *customerUsecase) DeleteCustomer(id int64) (err error) {
	return
}
func (m *customerUsecase) UpdateCustomer(customer *model.Customer) (err error) {
	return
}

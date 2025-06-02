package customer

import "pos/internal/model"

type CustomerUsecase interface {
	SaveCustomer(customer *model.Customer) (err error)
	GetAllCustomer(page, limit int64, search string) (customers []model.Customer, total int64, err error)
	GetCustomerDetails(id int64) (customer model.Customer, err error)
	DeleteCustomer(id int64) (err error)
	UpdateCustomer(customer *model.Customer) (err error)
}

type CustomerRepository interface {
	SaveCustomer(customer *model.Customer) (err error)
	GetAllCustomer(offset, limit int64, search string) (customers []model.Customer, total int64, err error)
	GetCustomerDetails(id int64) (customer model.Customer, err error)
	DeleteCustomer(id int64) (err error)
	UpdateCustomer(customer *model.Customer) (err error)
}

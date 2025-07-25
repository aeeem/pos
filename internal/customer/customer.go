package customer

import (
	"pos/internal/domain"
)

type CustomerUsecase interface {
	SaveCustomer(customer *domain.Customer) (err error)
	GetAllCustomer(page, limit int64, search string) (customers []domain.Customer, total int64, err error)
	GetCustomerDetails(id int64) (customer domain.Customer, err error)
	DeleteCustomer(id int64) (err error)
	UpdateCustomer(customer *domain.Customer) (err error)
}

type CustomerRepository interface {
	SaveCustomer(customer *domain.Customer) (err error)
	GetAllCustomer(offset, limit int64, search string) (customers []domain.Customer, total int64, err error)
	GetCustomerDetails(id int64) (customer domain.Customer, err error)
	DeleteCustomer(id int64) (err error)
	UpdateCustomer(customer *domain.Customer) (err error)
}

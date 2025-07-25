package mutation

import (
	"pos/internal/domain"
)

type MutationRepository interface {
	GetCustomerMutation(customerID uint, MutationType string) (CustomerMutation []domain.Mutation, err error)
	SaveCustomerMutation(mutation *domain.Mutation) (err error)
	UpdateCustomerMutation(mutationID uint, mutation *domain.Mutation) (err error)
	GetCustomerBalance(customerID uint) (CustomerBalance float64, err error)
}

type MutationUsecase interface {
	GetCustomerMutation(customerID uint, MutationType string) (CustomerMutation []domain.Mutation, err error)
	SaveCustomerMutation(mutation *domain.Mutation) (err error)
	UpdateCustomerMutation(mutationID uint, mutation *domain.Mutation) (err error)
	GetCustomerBalance(customerID uint) (CustomerBalance float64, err error)
}

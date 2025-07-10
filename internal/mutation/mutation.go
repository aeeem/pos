package mutation

import "pos/internal/model"

type MutationRepository interface {
	GetCustomerMutation(customerID uint, MutationType string) (CustomerMutation []model.Mutation, err error)
	SaveCustomerMutation(mutation *model.Mutation) (err error)
	UpdateCustomerMutation(mutationID uint, mutation *model.Mutation) (err error)
	GetCustomerBalance(customerID uint) (CustomerBalance float64, err error)
}

type MutationUsecase interface {
	GetCustomerMutation(customerID uint, MutationType string) (CustomerMutation []model.Mutation, err error)
	SaveCustomerMutation(mutation *model.Mutation) (err error)
	UpdateCustomerMutation(mutationID uint, mutation *model.Mutation) (err error)
	GetCustomerBalance(customerID uint) (CustomerBalance float64, err error)
}

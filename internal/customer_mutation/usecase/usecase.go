package usecase

import (
	customermutation "pos/internal/customer_mutation"
	"pos/internal/model"
)

type CustomerDebtMutationUsecase struct {
	DebtMutationRepository customermutation.CustomerMutationRepository
}

func NewCustomerDebtMutationUsecase(c customermutation.CustomerMutationRepository) customermutation.CustomerMutationUsecase {
	return &CustomerDebtMutationUsecase{
		DebtMutationRepository: c,
	}
}

func (Cm *CustomerDebtMutationUsecase) SaveCustomerMutation(DebtMutation *model.CustomerDebtMutations) (err error) {
	err = Cm.DebtMutationRepository.SaveCustomerMutation(DebtMutation)
	return
}

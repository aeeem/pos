package customermutation

import (
	"pos/internal/domain"
)

type CustomerMutationRepository interface {
	SaveCustomerMutation(mutation *domain.CustomerDebtMutations) (err error)
}

type CustomerMutationUsecase interface {
	SaveCustomerMutation(mutation *domain.CustomerDebtMutations) (err error)
}

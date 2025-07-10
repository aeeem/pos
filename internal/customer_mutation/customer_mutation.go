package customermutation

import (
	"pos/internal/model"
)

type CustomerMutationRepository interface {
	SaveCustomerMutation(mutation *model.CustomerDebtMutations) (err error)
}

type CustomerMutationUsecase interface {
	SaveCustomerMutation(mutation *model.CustomerDebtMutations) (err error)
}

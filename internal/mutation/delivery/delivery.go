package delivery

import (
	"pos/internal/mutation"
	"pos/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type MutationHandler struct {
	MutationUsecase mutation.MutationUsecase
	Validator       validator.XValidator
}

func NewMutationHandler(fiber *fiber.App, MutationUsecase mutation.MutationUsecase, validator *validator.XValidator) {
	MutationHanlder := MutationHandler{
		MutationUsecase: MutationUsecase,
		Validator:       *validator,
	}
	fiber.Get("/mutation/:transaction_id", MutationHanlder.GetCustomerMutation)
}

func (M *MutationHandler) GetCustomerMutation(c *fiber.Ctx) (err error) {
	return
}

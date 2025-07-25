package usecase

import (
	"pos/internal/domain"
	"pos/internal/mutation"
)

type MutationUsecase struct {
	MutationRepository mutation.MutationRepository
}

func NewMutationUsecase(MutationRepository mutation.MutationRepository) mutation.MutationUsecase {
	return &MutationUsecase{
		MutationRepository: MutationRepository,
	}
}

func (m *MutationUsecase) GetCustomerMutation(customerID uint, MutationType string) (CustomerMutation []domain.Mutation, err error) {
	CustomerMutation, err = m.MutationRepository.GetCustomerMutation(customerID, MutationType)
	return
}
func (m *MutationUsecase) SaveCustomerMutation(mutation *domain.Mutation) (err error) {
	err = m.MutationRepository.SaveCustomerMutation(mutation)
	return
}
func (m *MutationUsecase) UpdateCustomerMutation(mutationID uint, mutation *domain.Mutation) (err error) {
	err = m.MutationRepository.UpdateCustomerMutation(mutationID, mutation)
	return
}
func (m *MutationUsecase) GetCustomerBalance(customerID uint) (CustomerBalance float64, err error) {
	CustomerBalance, err = m.MutationRepository.GetCustomerBalance(customerID)
	return
}

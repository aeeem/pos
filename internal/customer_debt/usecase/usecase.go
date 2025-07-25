package usecase

import (
	"errors"
	"fmt"
	customerdebt "pos/internal/customer_debt"
	customermutation "pos/internal/customer_mutation"
	"pos/internal/domain"
	"pos/internal/mutation"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var msg = "Bulk payment customer debt"

type CustomerDebtUsecase struct {
	CustomerDebt        customerdebt.CustomerDebtRepository
	MutationUsecase     mutation.MutationUsecase
	MutationDebtUsecase customermutation.CustomerMutationUsecase
}

func NewCustomerUsecase(Cr customerdebt.CustomerDebtRepository,
	MutationUsecase mutation.MutationUsecase,
	MutationDebtUsecase customermutation.CustomerMutationUsecase) customerdebt.CustomerDebtUsecase {
	return &CustomerDebtUsecase{CustomerDebt: Cr,
		MutationUsecase:     MutationUsecase,
		MutationDebtUsecase: MutationDebtUsecase,
	}
}

func (c CustomerDebtUsecase) PayCustomerDebt(customerID int, transaction_id []int, amount float64) (Transactions []domain.Transaction, Debts []domain.CustomerDebt, err error) {
	//get latest customer mutation incase there's balance available
	//get last customer balance

	if amount < 0 {
		return []domain.Transaction{}, []domain.CustomerDebt{}, errors.New("amount must be greater than zero")
	} else if amount == 0 {
		msg = "payment using customer balance"
	}

	Balance, err := c.MutationUsecase.GetCustomerBalance(uint(customerID))
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return
	}
	BalRef := Balance
	Balance += amount

	TotalDebt := float64(0)
	CustomerDebts := []domain.CustomerDebt{}
	for _, v := range transaction_id {
		//get customer debt
		thisTransactionDebt, err := c.CustomerDebt.GetByTransactionID(uint(v))
		if err != nil {
			return []domain.Transaction{}, []domain.CustomerDebt{}, err
		}
		TotalDebt += thisTransactionDebt.UnpaidAmount
		ThisCurrentDebt := thisTransactionDebt.UnpaidAmount
		log.Info().Float64("totaldebt", TotalDebt).Msg("totaldebt")
		if ThisCurrentDebt <= Balance { //check if customer debt is less than Balance
			Balance = Balance - ThisCurrentDebt //reduce Balance
			thisTransactionDebt.PaidAmount = ThisCurrentDebt
			thisTransactionDebt.UnpaidAmount = 0 // because already paid
			err = c.CustomerDebt.PayCustomerDebt(customerID, "paid", v, ThisCurrentDebt)
			if err != nil {
				return []domain.Transaction{}, []domain.CustomerDebt{}, err
			}
			CustomerDebts = append(CustomerDebts, thisTransactionDebt)
		} else if Balance <= 0 {
			err = errors.New("not enough balance")
			return []domain.Transaction{}, []domain.CustomerDebt{}, err
		} else if ThisCurrentDebt > Balance { //if debt is more than Balance
			//set half paid
			ThisCurrentDebt = ThisCurrentDebt - Balance
			thisTransactionDebt.PaidAmount = Balance
			thisTransactionDebt.UnpaidAmount = ThisCurrentDebt
			// set it to 0 because the money already half paid the debt
			err = c.CustomerDebt.PayCustomerDebt(customerID, "half_paid", v, Balance)
			if err != nil {
				return []domain.Transaction{}, []domain.CustomerDebt{}, err
			}
			Balance = 0
			CustomerDebts = append(CustomerDebts, thisTransactionDebt)
			//update transactions

			break //set break because if Balance is less than debt no need to check another one

		}
	}
	//Insert customer mutation with latest balance which was Balance
	//create mutation with latest balance

	mutation := domain.Mutation{
		CustomerID:      uint(customerID),
		MutationType:    "cash_in",
		Amount:          amount, //keep track latest payment amount
		CustomerBalance: Balance,
		UsedAmount:      TotalDebt,
		MutationUUID:    fmt.Sprintf("%s-%v-%v-%v", "cash_in", transaction_id, customerID, time.Now().UnixNano()),
		Description:     msg,
	}
	if BalRef < TotalDebt && amount < TotalDebt {
		mutation.CustomerBalance = 0
		mutation.UsedAmount = BalRef
		mutation.Amount = amount
	}
	// if mutation.CustomerBalance <= mutation.UsedAmount && customer {
	// 	mutation.UsedAmount = mutation.CustomerBalance
	// }
	log.Info().Any("mutation", mutation).Msg("mutation")
	err = c.MutationUsecase.SaveCustomerMutation(&mutation)
	if err != nil {
		return
	}
	for _, v := range CustomerDebts {
		DebtMutation := domain.CustomerDebtMutations{
			CustomerDebtID: v.ID,
			MutationID:     mutation.ID,
			CurrentDebt:    v.UnpaidAmount,
			CurrentPayment: v.PaidAmount,
			TotalPayment:   mutation.Amount,
		}
		err = c.MutationDebtUsecase.SaveCustomerMutation(&DebtMutation)
		if err != nil {
			return
		}
	}

	return
}

func (C CustomerDebtUsecase) GetByCustomerID(customerID uint, debtType string, dateFrom string, dateTo string, page, limit int, search string) (res []domain.CustomerDebt, total int64, err error) {
	res, total, err = C.CustomerDebt.GetByCustomerID(customerID, debtType, dateFrom, dateTo, page, limit, search)
	if err != nil {
		return
	}
	return
}
func (C CustomerDebtUsecase) Fetch(debtType string, dateFrom string, dateTo string, page, limit int, search string) (res []domain.CustomerDebt, total int64, err error) {
	res, total, err = C.CustomerDebt.Fetch(debtType, dateFrom, dateTo, page, limit, search)
	if err != nil {
		return
	}
	return
}

func (C CustomerDebtUsecase) GetDebtDetails(debtID uint) (res domain.CustomerDebt, err error) {
	res, err = C.CustomerDebt.GetDebtDetails(debtID)
	if err != nil {
		return
	}

	return
}

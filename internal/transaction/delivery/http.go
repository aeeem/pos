package delivery

import (
	"encoding/json"
	"fmt"
	"pos/internal/helper"
	"pos/internal/http_error"
	"pos/internal/model"
	"pos/internal/transaction"
	"pos/internal/validator"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	Validator          *validator.XValidator
	TransactionUsecase transaction.TransactionUsecase
}

func NewTransactionHandler(transactionUsecase transaction.TransactionUsecase, f *fiber.App, validator *validator.XValidator) {
	TransactionHandler := TransactionHandler{
		TransactionUsecase: transactionUsecase,
		Validator:          validator,
	}
	f.Get("/transaction", TransactionHandler.GetTransactions)
	f.Get("/transaction/:id", TransactionHandler.GetTransactionDetails)
	f.Post("/transaction", TransactionHandler.Savetransaction)

	f.Put("/transaction", TransactionHandler.Updatetransaction)
	f.Delete("/transaction/:id", TransactionHandler.Deletetransaction)
}

func (t TransactionHandler) GetTransactions(c *fiber.Ctx) error {
	GetRequest := GetTransactionsRequest{
		GetRequest: helper.GetRequest{
			Page:   c.QueryInt("page"),
			Limit:  c.QueryInt("limit"),
			Search: c.Query("search"),
		},
		Status:     c.Query("status"),
		CustomerID: c.QueryInt("customer_id"),
	}
	if errs := t.Validator.Validate(GetRequest); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: strings.Join(errMsgs, " and "),
		}
	}
	transactions, total, err := t.TransactionUsecase.GetTransactions(int64(GetRequest.Page), int64(GetRequest.Limit), GetRequest.Search, model.Status(GetRequest.Status), int64(GetRequest.CustomerID))
	if err != nil {
		herr := http_error.CheckError(err)
		return c.JSON(fiber.Error{
			Code:    herr.HTTPErrorCode,
			Message: herr.Message,
		})
	}
	return c.JSON(
		GetItemResponse{
			Transaction: transactions,
			Total:       total,
			Page:        GetRequest.Page,
			Limit:       GetRequest.Limit,
		},
	)
}

func (t TransactionHandler) GetTransactionDetails(c *fiber.Ctx) error {
	return nil
}
func (t TransactionHandler) Savetransaction(c *fiber.Ctx) error {
	SaveItemRequest := SaveOrUpdate{}
	if err := json.Unmarshal(c.Body(), &SaveItemRequest); err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}

	if errs := t.Validator.Validate(SaveItemRequest); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: strings.Join(errMsgs, " and "),
		}
	}
	Transaction := model.Transaction{
		CustomerName: SaveItemRequest.CustomerName,
	}
	if SaveItemRequest.CustomerID != 0 {
		Transaction.CustomerID = uint(SaveItemRequest.CustomerID)

	}
	if SaveItemRequest.Status != nil {
		Transaction.Status = model.Status(*SaveItemRequest.Status)
	}

	if len(SaveItemRequest.Cart) > 0 {
		for _, v := range SaveItemRequest.Cart {
			Transaction.Cart = append(Transaction.Cart, model.Cart{
				TransactionID: uint(v.TransactionID),
				ItemID:        uint(v.ItemID),
				Quantity:      v.Quantity,
				PriceID:       uint(v.PriceID),
				SubPrice:      0,
			})
		}
	}
	err := t.TransactionUsecase.Savetransaction(&Transaction)
	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}
	return c.JSON(Transaction)
}

func (t TransactionHandler) Updatetransaction(c *fiber.Ctx) error {
	return nil
}
func (t TransactionHandler) Deletetransaction(c *fiber.Ctx) error {
	return nil
}

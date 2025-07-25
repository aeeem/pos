package delivery

import (
	"encoding/json"
	"fmt"
	"pos/internal/domain"
	"pos/internal/helper"
	"pos/internal/http_error"
	"pos/internal/transaction"
	"pos/internal/validator"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
	f.Put("/transaction/:id", TransactionHandler.Updatetransaction)
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
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: strings.Join(errMsgs, " and "),
		})
	}
	transactions, total, err := t.TransactionUsecase.GetTransactions(int64(GetRequest.Page), int64(GetRequest.Limit), GetRequest.Search, domain.Status(GetRequest.Status), int64(GetRequest.CustomerID))
	if err != nil {
		herr := http_error.CheckError(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Error{
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
	//get item id
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: "id should be a number",
		})
	}
	if id == 0 {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: "id is required",
		})
	}

	transaction, err := t.TransactionUsecase.GetTransactionDetails(int64(id))
	if err != nil {
		herr := http_error.CheckError(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Error{
			Code:    herr.HTTPErrorCode,
			Message: herr.Message,
		})
	}
	return c.JSON(helper.StandardResponse{
		Status:  "success",
		Message: "transaction details",
		Data:    transaction,
	})
}

func (t TransactionHandler) Savetransaction(c *fiber.Ctx) error {
	SaveItemRequest := SaveOrUpdate{}
	if err := json.Unmarshal(c.Body(), &SaveItemRequest); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		})
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
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: strings.Join(errMsgs, " and "),
		})
	}
	Transaction := domain.Transaction{
		CustomerName: SaveItemRequest.CustomerName,
	}
	if SaveItemRequest.CustomerID != 0 {
		Transaction.CustomerID = uint(SaveItemRequest.CustomerID)

	}
	if SaveItemRequest.Status != nil {
		Transaction.Status = domain.Status(*SaveItemRequest.Status)
	}

	if SaveItemRequest.CustomerTransactionNo != 0 {
		Transaction.CustomerTransactionNo = uint(SaveItemRequest.CustomerTransactionNo)
	}
	if len(SaveItemRequest.Cart) > 0 {
		for _, v := range SaveItemRequest.Cart {
			Transaction.Cart = append(Transaction.Cart, domain.Cart{
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
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		})
	}
	return c.JSON(Transaction)
}

func (t TransactionHandler) Updatetransaction(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	SaveItemRequest := SaveOrUpdate{}
	if err := json.Unmarshal(c.Body(), &SaveItemRequest); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		})
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
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: strings.Join(errMsgs, " and "),
		})
	}
	Transaction := domain.Transaction{
		Model: gorm.Model{
			ID: uint(id),
		},
		CustomerName: SaveItemRequest.CustomerName,
	}
	if SaveItemRequest.CustomerID != 0 {
		Transaction.CustomerID = uint(SaveItemRequest.CustomerID)

	}
	if SaveItemRequest.Status != nil {
		Transaction.Status = domain.Status(*SaveItemRequest.Status)
	}
	if SaveItemRequest.CustomerTransactionNo != 0 {
		Transaction.CustomerTransactionNo = uint(SaveItemRequest.CustomerTransactionNo)
	}
	err := t.TransactionUsecase.UpdateTransaction(&Transaction)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		})
	}

	return c.JSON(helper.StandardResponse{
		Status:  "success",
		Message: "transaction Updated successfully",
		Data:    Transaction,
	})
}
func (t TransactionHandler) Deletetransaction(c *fiber.Ctx) error {
	//get item id
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: "id should be a number",
		})
	}
	if id == 0 {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: "id is required",
		})
	}
	err = t.TransactionUsecase.DeleteTransaction(int64(id))
	if err != nil {
		herr := http_error.CheckError(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Error{
			Code:    herr.HTTPErrorCode,
			Message: herr.Message,
		})
	}
	return c.JSON(helper.StandardResponse{
		Status:  "success",
		Message: "transaction deleted successfully",
		Data:    nil,
	})
}

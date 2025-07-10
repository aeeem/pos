package delivery

import (
	"encoding/json"
	"fmt"
	customerdebt "pos/internal/customer_debt"
	"pos/internal/helper"
	"pos/internal/http_error"
	"pos/internal/validator"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type DebtHandler struct {
	Validator   *validator.XValidator
	DebtUsecase customerdebt.CustomerDebtUsecase
}

func NewCustomerdebtHandler(f *fiber.App, validator *validator.XValidator, DebtUsecase customerdebt.CustomerDebtUsecase) {
	DebtHandler := DebtHandler{
		DebtUsecase: DebtUsecase,
		Validator:   validator,
	}
	f.Get("/customer_debt", DebtHandler.GetDebt)
	f.Get("/customer_debt/customer/:id", DebtHandler.GetDebtByCustomerID)
	f.Get("/customer_debt/mutation/:id", DebtHandler.MutationDebt)

	f.Post("/customer_debt/:cust_id/pay", DebtHandler.PayCustomerDebt)
}

func (D *DebtHandler) MutationDebt(c *fiber.Ctx) (err error) {
	DebtID, err := c.ParamsInt("id")
	if err != nil {
		log.Info().Err(err).Msg("error occured")
		return c.Status(400).JSON(&fiber.Error{
			Code:    400,
			Message: "id should be a number",
		})
	}

	res, err := D.DebtUsecase.GetDebtDetails(uint(DebtID))
	if err != nil {
		herr := http_error.CheckError(err)
		return c.JSON(fiber.Error{
			Code:    herr.HTTPErrorCode,
			Message: herr.Message,
		})
	}

	return c.JSON(helper.StandardResponse{
		Status: "success",
		Data:   res,
	})
}

func (D *DebtHandler) GetDebt(c *fiber.Ctx) (err error) {
	request := Fetch{}
	request.Limit = c.QueryInt("limit")
	request.Page = c.QueryInt("page")

	if c.Query("date_to") != "" {
		request.DateTo = c.Query("date_to")
	}
	if c.Query("date_from") != "" {
		request.DateFrom = c.Query("date_from")
	}
	if c.Query("debt_type") != "" {
		request.DebtType = c.Query("debt_type")
	}
	if c.Query("search") != "" {
		request.Search = c.Query("search")
	}
	if errs := D.Validator.Validate(request); len(errs) > 0 && errs[0].Error {
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

	res, total, err := D.DebtUsecase.Fetch(request.DebtType,
		request.DateFrom,
		request.DateTo,
		int(helper.PageToOffset(int64(request.Page), int64(request.Limit))),
		request.Limit,
		request.Search)
	if err != nil {
		herr := http_error.CheckError(err)
		return c.JSON(fiber.Error{
			Code:    herr.HTTPErrorCode,
			Message: herr.Message,
		})
	}

	return c.JSON(GetResponse{
		ListResponse: helper.ListResponse{
			Total: total,
			Page:  int64(request.Page),
			Limit: int64(request.Limit),
		},
		Data: res,
	})
}

func (D *DebtHandler) GetDebtByCustomerID(c *fiber.Ctx) (err error) {

	request := GetCustomeDebtRequest{}
	request.CustomerID, _ = c.ParamsInt("id")
	request.Limit = c.QueryInt("limit")
	request.Page = c.QueryInt("page")

	if c.Query("date_to") != "" {
		request.DateTo = c.Query("date_to")
	}
	if c.Query("date_from") != "" {
		request.DateFrom = c.Query("date_from")
	}
	if c.Query("debt_type") != "" {
		request.DebtType = c.Query("debt_type")
	}
	if c.Query("search") != "" {
		request.Search = c.Query("search")
	}
	if errs := D.Validator.Validate(request); len(errs) > 0 && errs[0].Error {
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

	res, total, err := D.DebtUsecase.GetByCustomerID(
		uint(request.CustomerID),
		request.DebtType,
		request.DateFrom,
		request.DateTo,
		int(helper.PageToOffset(int64(request.Page), int64(request.Limit))),
		request.Limit,
		request.Search)
	if err != nil {
		herr := http_error.CheckError(err)
		return c.JSON(fiber.Error{
			Code:    herr.HTTPErrorCode,
			Message: herr.Message,
		})
	}

	return c.JSON(GetResponse{
		ListResponse: helper.ListResponse{
			Total: total,
			Page:  int64(request.Page),
			Limit: int64(request.Limit),
		},
		Data: res,
	})
}

func (D DebtHandler) PayCustomerDebt(c *fiber.Ctx) (err error) {
	CustomerID, err := c.ParamsInt("cust_id")
	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}

	PayRequest := PayCustomerDebtRequest{}
	err = json.Unmarshal(c.Body(), &PayRequest)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		})
	}
	if errs := D.Validator.Validate(PayRequest); len(errs) > 0 && errs[0].Error {
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

	TransactionList, debts, err := D.DebtUsecase.PayCustomerDebt(CustomerID, PayRequest.TransactionIDs, PayRequest.TotalAmount)
	if err != nil {
		herr := http_error.CheckError(err)
		return c.JSON(fiber.Error{
			Code:    herr.HTTPErrorCode,
			Message: herr.Message,
		})
	}
	PayResponse := PayDebtResponse{
		Transactions: TransactionList,
		Debt:         debts,
	}
	return c.JSON(
		helper.StandardResponse{
			Status:  "success",
			Message: "debt paid successfully",
			Data:    PayResponse,
		},
	)
}

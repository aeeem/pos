package delivery

import (
	"encoding/json"
	"fmt"
	"pos/internal/customer"
	"pos/internal/domain"
	"pos/internal/helper"
	"pos/internal/http_error"
	"pos/internal/validator"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	CustomerUsecase customer.CustomerUsecase
	Validator       validator.XValidator
}

func NewCustomerHandler(fiber *fiber.App, customerUsecase customer.CustomerUsecase, validator *validator.XValidator) {
	CustomerHandler := CustomerHandler{
		CustomerUsecase: customerUsecase,
		Validator:       *validator,
	}
	fiber.Get("/customer", CustomerHandler.GetCustomers)
	fiber.Post("/customer", CustomerHandler.SaveCustomer)
}

func (h *CustomerHandler) GetCustomers(c *fiber.Ctx) (err error) {
	GetRequest := helper.GetRequest{

		Page:   c.QueryInt("page"),
		Limit:  c.QueryInt("limit"),
		Search: c.Query("search"),
	}
	if errs := h.Validator.Validate(GetRequest); len(errs) > 0 && errs[0].Error {
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

	customers, total, err := h.CustomerUsecase.GetAllCustomer(int64(GetRequest.Page), int64(GetRequest.Limit), GetRequest.Search)
	if err != nil {
		herr := http_error.CheckError(err)
		return c.JSON(fiber.Error{
			Code:    herr.HTTPErrorCode,
			Message: herr.Message,
		})
	}
	return c.JSON(CustomerListResponse{
		ListResponse: helper.ListResponse{
			Total: total,
			Page:  int64(GetRequest.Page),
			Limit: int64(GetRequest.Limit),
		},
		Customers: customers,
	})
}

func (h *CustomerHandler) SaveCustomer(c *fiber.Ctx) (err error) {

	SaveRequest := SaveOrUpdate{}
	if err := json.Unmarshal(c.Body(), &SaveRequest); err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}
	if errs := h.Validator.Validate(SaveRequest); len(errs) > 0 && errs[0].Error {
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
	Customer := domain.Customer{
		CustomerName: SaveRequest.CustomerName,
		PhoneNumber:  SaveRequest.PhoneNumber,
	}
	err = h.CustomerUsecase.SaveCustomer(&Customer)
	if err != nil {
		herr := http_error.CheckError(err)
		return c.Status(herr.HTTPErrorCode).JSON(fiber.Error{
			Code:    herr.HTTPErrorCode,
			Message: herr.Message,
		})
	}
	return c.Status(fiber.StatusOK).JSON(
		CustomerSaveOrUpdateResponse{
			Customer: Customer,
			StandardResponse: helper.StandardResponse{
				Status:  "success",
				Message: "Customer saved successfully",
			},
		},
	)
}

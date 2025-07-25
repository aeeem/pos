package delivery

import (
	"encoding/json"
	"fmt"
	"pos/internal/cart"
	"pos/internal/domain"
	"pos/internal/helper"
	"pos/internal/http_error"
	"pos/internal/validator"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CartHandler struct {
	Validator   *validator.XValidator
	CartUsecase cart.CartUsecase
}

func NewCartHandler(f *fiber.App, validator *validator.XValidator, CartUscase cart.CartUsecase) {
	CartHandler := CartHandler{
		CartUsecase: CartUscase,
		Validator:   validator,
	}
	f.Get("/Cart", CartHandler.GetCartByTransactionID)
	f.Post("/Cart", CartHandler.SaveCart)
	f.Delete("/Cart/:id", CartHandler.DeleteCart)
	f.Put("/Cart", CartHandler.UpdateCart)
}

func (m CartHandler) GetCartByTransactionID(c *fiber.Ctx) (err error) {
	GetRequest := GetCartByTransactionIDRequest{
		GetRequest: helper.GetRequest{
			Page:   c.QueryInt("page"),
			Limit:  c.QueryInt("limit"),
			Search: c.Query("search"),
		},
		TransactionID: int64(c.QueryInt("transaction_id")),
	}
	if errs := m.Validator.Validate(GetRequest); len(errs) > 0 && errs[0].Error {
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

	cart, total, err := m.CartUsecase.GetCartByTransactionID(int64(GetRequest.Page), int64(GetRequest.Limit), uint(GetRequest.TransactionID))
	if err != nil {
		herr := http_error.CheckError(err)
		return c.JSON(fiber.Error{
			Code:    herr.HTTPErrorCode,
			Message: herr.Message,
		})
	}
	return c.JSON(GetCartByTransactionIDResponse{
		ListResponse: helper.ListResponse{
			Total: total,
			Page:  int64(GetRequest.Page),
			Limit: int64(GetRequest.Limit),
		},
		Carts: cart,
	})
}

func (m CartHandler) SaveCart(c *fiber.Ctx) (err error) {
	SaveItemRequest := SaveOrUpdate{}
	if err := json.Unmarshal(c.Body(), &SaveItemRequest); err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}
	if errs := m.Validator.Validate(SaveItemRequest); len(errs) > 0 && errs[0].Error {
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
	cart := domain.Cart{
		TransactionID: uint(SaveItemRequest.TransactionID),
		ItemID:        uint(SaveItemRequest.ItemID),
		Quantity:      SaveItemRequest.Quantity,
		PriceID:       uint(SaveItemRequest.PriceID),
		SubPrice:      0,
	}
	err = m.CartUsecase.SaveCart(&cart)
	if err != nil {
		herr := http_error.CheckError(err)
		return c.JSON(fiber.Error{
			Code:    herr.HTTPErrorCode,
			Message: herr.Message,
		})
	}
	return c.JSON(
		SaveCartResponse{
			Status: "success",
			Cart:   cart,
		},
	)
}

func (m CartHandler) UpdateCart(c *fiber.Ctx) (err error) {
	SaveItemRequest := SaveOrUpdate{}
	if err := json.Unmarshal(c.Body(), &SaveItemRequest); err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}

	if errs := m.Validator.Validate(SaveItemRequest); len(errs) > 0 && errs[0].Error {
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
	cart := domain.Cart{
		Model: gorm.Model{
			ID: uint(SaveItemRequest.ID),
		},
		TransactionID: uint(SaveItemRequest.TransactionID),
		ItemID:        uint(SaveItemRequest.ItemID),
		Quantity:      SaveItemRequest.Quantity,
		PriceID:       uint(SaveItemRequest.PriceID),
		SubPrice:      0,
	}
	err = m.CartUsecase.SaveCart(&cart)
	if err != nil {
		herr := http_error.CheckError(err)
		return c.JSON(fiber.Error{
			Code:    herr.HTTPErrorCode,
			Message: herr.Message,
		})
	}
	return c.JSON(
		SaveCartResponse{
			Status: "success",
			Cart:   cart,
		},
	)
}

func (m CartHandler) DeleteCart(c *fiber.Ctx) (err error) {
	print("sampe")
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
	err = m.CartUsecase.DeleteCart(uint(id))
	if err != nil {
		herr := http_error.CheckError(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Error{
			Code:    herr.HTTPErrorCode,
			Message: herr.Message,
		})
	}
	return c.JSON(helper.StandardResponse{
		Status:  "success",
		Message: "cart deleted successfully",
		Data:    nil,
	})
}

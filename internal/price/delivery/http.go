package delivery

import (
	"encoding/json"
	"fmt"
	"pos/internal/helper"
	"pos/internal/http_error"
	"pos/internal/model"
	"pos/internal/price"
	"pos/internal/validator"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type PriceHandler struct {
	Validator    *validator.XValidator
	PriceUsecase price.PriceUsecase
}

func NewPriceHandler(f *fiber.App, validator *validator.XValidator, priceUscase price.PriceUsecase) {
	PriceHandler := PriceHandler{
		PriceUsecase: priceUscase,
		Validator:    validator,
	}
	f.Get("/price", PriceHandler.GetPrices)
	f.Get("/price/:id", PriceHandler.GetPriceDetails)
	f.Post("/price", PriceHandler.SavePrice)
	f.Put("/price", PriceHandler.UpdatePrice)
	f.Delete("/price/:id", PriceHandler.DeletePrice)
}

func (h *PriceHandler) GetPrices(c *fiber.Ctx) error {

	GetRequest := GetPriceRequest{
		GetRequest: helper.GetRequest{
			Page:   c.QueryInt("page"),
			Limit:  c.QueryInt("limit"),
			Search: c.Query("search"),
		},
		ItemID: int64(c.QueryInt("item_id")),
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
	prices, total, err := h.PriceUsecase.GetPrices(int64(GetRequest.Page), int64(GetRequest.Limit), GetRequest.Search, GetRequest.ItemID)
	if err != nil {
		herr := http_error.CheckError(err)
		return c.JSON(fiber.Error{
			Code:    herr.HTTPErrorCode,
			Message: herr.Message,
		})
	}

	return c.JSON(GetItemResponse{
		Price: prices,
		Total: total,
		Page:  GetRequest.Page,
		Limit: GetRequest.Limit,
	})

}

func (h *PriceHandler) GetPriceDetails(c *fiber.Ctx) error {
	return nil
}

func (h *PriceHandler) SavePrice(c *fiber.Ctx) error {
	SaveItemRequest := SaveOrUpdate{}
	if err := json.Unmarshal(c.Body(), &SaveItemRequest); err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}

	if errs := h.Validator.Validate(SaveItemRequest); len(errs) > 0 && errs[0].Error {
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
	Price := model.Price{
		Price:  SaveItemRequest.Price,
		Active: SaveItemRequest.Active,
		ItemID: uint(SaveItemRequest.ItemID),
	}
	err := h.PriceUsecase.SavePrice(&Price)
	if err != nil {
		herr := http_error.CheckError(err)
		return c.JSON(fiber.Error{
			Code:    herr.HTTPErrorCode,
			Message: herr.Message,
		})
	}
	return c.JSON(
		SavePriceResponse{
			Status: "success",
			Price:  Price,
		},
	)
}

func (h *PriceHandler) UpdatePrice(c *fiber.Ctx) error {
	return nil
}

func (h *PriceHandler) DeletePrice(c *fiber.Ctx) error {
	return nil
}

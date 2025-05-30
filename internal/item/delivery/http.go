package delivery

import (
	"encoding/json"
	"fmt"
	"pos/internal/helper"
	http_error "pos/internal/http_error"
	"pos/internal/item"
	"pos/internal/model"
	"pos/internal/validator"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type ItemHandler struct {
	Validator   *validator.XValidator
	ItemUsecase item.ItemUsecase
}

func NewItemHandler(f *fiber.App, validator *validator.XValidator, itemUsecase item.ItemUsecase) {
	ItemHandler := ItemHandler{
		ItemUsecase: itemUsecase,
		Validator:   validator,
	}
	f.Get("/item", ItemHandler.GetItems)
	f.Get("/item/:id", ItemHandler.GetItemDetails)
	f.Post("/item", ItemHandler.SaveItem)
	f.Put("/item", ItemHandler.UpdateItem)
	f.Delete("/item/:id", ItemHandler.DeleteItem)
}

func (h *ItemHandler) GetItems(c *fiber.Ctx) error {
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
	items, total, err := h.ItemUsecase.GetItems(int64(GetRequest.Page), int64(GetRequest.Limit), GetRequest.Search)
	if err != nil {
		herr := http_error.CheckError(err)
		return c.JSON(fiber.Error{
			Code:    herr.HTTPErrorCode,
			Message: herr.Message,
		})
	}
	return c.JSON(
		GetItemResponse{
			Items: items,
			Total: total,
			Page:  GetRequest.Page,
			Limit: GetRequest.Limit,
		},
	)
}
func (h *ItemHandler) GetItemDetails(c *fiber.Ctx) error {
	return nil
}

func (h *ItemHandler) SaveItem(c *fiber.Ctx) error {
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
	Items := model.Item{
		ItemName:     SaveItemRequest.ItemName,
		MaxPriceItem: SaveItemRequest.MaxPrice,
	}
	err := h.ItemUsecase.SaveItem(&Items)
	if err != nil {
		herr := http_error.CheckError(err)
		return c.JSON(fiber.Error{
			Code:    herr.HTTPErrorCode,
			Message: herr.Message,
		})
	}
	return c.JSON(
		SaveItemResponse{
			Status: "success",
			Item:   Items,
		},
	)
}
func (h *ItemHandler) UpdateItem(c *fiber.Ctx) error {
	return nil
}
func (h *ItemHandler) DeleteItem(c *fiber.Ctx) error {
	return nil
}

package delivery

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"pos/internal/helper"
	http_error "pos/internal/http_error"
	"pos/internal/item"
	"pos/internal/model"
	"pos/internal/validator"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	f.Get("/images/:id", ItemHandler.Images)

	f.Get("/item", ItemHandler.GetItems)
	f.Get("/item/:id", ItemHandler.GetItemDetails)
	f.Post("/item", ItemHandler.SaveItem)
	f.Put("/item/:id", ItemHandler.UpdateItem)
	f.Delete("/item/:id", ItemHandler.DeleteItem)
}
func (h *ItemHandler) Images(c *fiber.Ctx) error {
	print("sampe")
	id := c.Params("id")
	fileBytes, err := os.ReadFile("images/" + id)
	if err != nil {
		return c.Status(404).JSON(&fiber.Error{
			Code:    fiber.ErrNotFound.Code,
			Message: "file not found",
		})
	}
	c.Set("Content-Type", "image/jpeg")
	return c.Send(fileBytes)
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

	item, err := h.ItemUsecase.GetItemDetails(int64(id))
	if err != nil {
		herr := http_error.CheckError(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Error{
			Code:    herr.HTTPErrorCode,
			Message: herr.Message,
		})
	}
	return c.JSON(helper.StandardResponse{
		Status:  "success",
		Message: "item details",
		Data:    item,
	})
}

func (h *ItemHandler) SaveItem(c *fiber.Ctx) error {
	//parse value as json
	formJson := c.FormValue("json-request")

	SaveItemRequest := SaveOrUpdate{}
	if err := json.Unmarshal([]byte(formJson), &SaveItemRequest); err != nil {
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
	//upload image
	file, err := c.FormFile("image")
	if err != nil {
		log.Print(err)
	} else {

		filename := fmt.Sprintf("%s-%s", uuid.New().String(), SaveItemRequest.ItemName)
		fileExt := strings.Split(file.Filename, ".")[1]
		filename = fmt.Sprintf("%s.%s", filename, fileExt)
		err = c.SaveFile(file, fmt.Sprintf("./images/%s", filename))
		if err != nil {
			return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": err.Error()})
		}
		imageUrl := fmt.Sprintf("http://47.236.241.247/images/%s", filename)

		Items.ImageUrl = imageUrl
	}

	for _, v := range SaveItemRequest.Price {
		Items.Price = append(Items.Price, model.Price{
			Price:  v.Price,
			Active: v.Active,
			Unit:   v.Unit,
		})

	}
	err = h.ItemUsecase.SaveItem(&Items)
	if err != nil {

		herr := http_error.CheckError(err)
		return c.JSON(fiber.Error{
			Code:    herr.HTTPErrorCode,
			Message: herr.Message,
		})
	}

	if err != nil {
		log.Println("Error in saving Image :", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	return c.JSON(
		SaveItemResponse{
			Status: "success",
			Item:   Items,
		},
	)
}
func (h *ItemHandler) UpdateItem(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	SaveItemRequest := SaveOrUpdate{}
	if err := json.Unmarshal(c.Body(), &SaveItemRequest); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		})
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
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: strings.Join(errMsgs, " and "),
		})
	}
	Item := model.Item{
		Model: gorm.Model{
			ID: uint(id),
		},
		ItemName: SaveItemRequest.ItemName,
	}
	if len(SaveItemRequest.Price) != 0 {
		for _, v := range SaveItemRequest.Price {
			Item.Price = append(Item.Price, model.Price{
				Model: gorm.Model{
					ID: uint(v.ID),
				},
				Price:  v.Price,
				Active: v.Active,
				ItemID: uint(id),
				Unit:   v.Unit,
			})

		}

	}

	err := h.ItemUsecase.UpdateItem(&Item)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		})
	}

	return c.JSON(helper.StandardResponse{
		Status:  "success",
		Message: "price updated successfully",
		Data:    Item,
	})
}
func (h *ItemHandler) DeleteItem(c *fiber.Ctx) error {
	return nil
}

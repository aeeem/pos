package http

import (
	"fmt"
	"log"
	itemHandler "pos/internal/item/delivery"
	itemrepository "pos/internal/item/repository"
	itemUsecase "pos/internal/item/usecase"

	internalValidator "pos/internal/validator"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var validate = validator.New()

func HttpRun(port string) {
	myValidator := &internalValidator.XValidator{
		Validator: validate,
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			viper.GetString("database.host"),
			viper.GetString("database.user"),
			viper.GetString("database.pass"),
			viper.GetString("database.name"), viper.GetString("database.port")),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	itemRepo := itemrepository.NewItemPresistenRepository(db)
	itemUC := itemUsecase.NewItemUsecase(itemRepo)
	itemHandler.NewItemHandler(app, myValidator, itemUC)
	log.Println(app.Listen(":" + port))
}

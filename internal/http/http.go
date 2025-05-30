package http

import (
	"fmt"
	"log"
	itemHandler "pos/internal/item/delivery"
	itemrepository "pos/internal/item/repository"
	"pos/internal/seeder"

	itemUsecase "pos/internal/item/usecase"
	"pos/internal/model"

	priceHandler "pos/internal/price/delivery"
	priceRepository "pos/internal/price/repository"
	priceUsecase "pos/internal/price/usecase"

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

	//migration
	db.AutoMigrate(&model.Item{}, &model.Price{}, &model.Transaction{}, &model.Cart{})

	if viper.GetString("seed") == "true" {
		seeder.SeedItem(db)
	}
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	itemRepo := itemrepository.NewItemPresistenRepository(db)
	itemUC := itemUsecase.NewItemUsecase(itemRepo)
	itemHandler.NewItemHandler(app, myValidator, itemUC)

	priceRepo := priceRepository.NewPricePresistentRepository(db)
	priceUC := priceUsecase.NewPriceUsecase(priceRepo)
	priceHandler.NewPriceHandler(app, myValidator, priceUC)

	log.Println(app.Listen(port))
}

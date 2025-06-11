package http

import (
	"fmt"
	"pos/internal/helper"
	itemHandler "pos/internal/item/delivery"
	itemrepository "pos/internal/item/repository"
	itemUsecase "pos/internal/item/usecase"

	"pos/internal/model"
	"pos/internal/seeder"

	transactionHandler "pos/internal/transaction/delivery"
	transactionRepository "pos/internal/transaction/repository"
	transactionUsecase "pos/internal/transaction/usecase"

	customerHandler "pos/internal/customer/delivery"
	customerRepository "pos/internal/customer/repository"
	customerUsecase "pos/internal/customer/usecase"

	CartHandler "pos/internal/cart/delivery"
	cartRepository "pos/internal/cart/repository"
	cartUsecase "pos/internal/cart/usecase"

	priceHandler "pos/internal/price/delivery"
	priceRepository "pos/internal/price/repository"
	priceUsecase "pos/internal/price/usecase"

	internalValidator "pos/internal/validator"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"

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
	//create enum
	helper.CreateStatusEnum(db)
	//migration
	db.AutoMigrate(&model.Item{}, &model.Price{}, &model.Transaction{}, &model.Cart{}, &model.Customer{})
	//creating trigger function after migrations
	helper.CheckCustomer(db)
	helper.CheckCustomerCountAfterUpdate(db)
	helper.UpdateTotalPrice(db)
	//creating table trigger after migrations
	helper.UpdateTransactionTrigger(db)
	helper.CartTrigger(db)
	helper.TransactionTrigger(db)

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

	cartRepository := cartRepository.NewcartPresistentRepository(db)
	cartUsecase := cartUsecase.NewCartUsecase(cartRepository, itemRepo, priceRepo)
	CartHandler.NewCartHandler(app, myValidator, cartUsecase)

	transactionRepo := transactionRepository.NewTransactionPresistentRepository(db)
	transactionUsecase := transactionUsecase.NewTransactionUsecase(transactionRepo, cartUsecase, itemUC, priceUC)
	transactionHandler.NewTransactionHandler(transactionUsecase, app, myValidator)

	customerRepo := customerRepository.NewCustomerPresistentRepository(db)
	customerUsecase := customerUsecase.NewCustomerUsecase(customerRepo)
	customerHandler.NewCustomerHandler(app, customerUsecase, myValidator)

	if viper.GetString("seed") == "true" {
		seeder.SeedItem(db)
		log.Info().Msg("item seeder")
		seeder.TransactionSeeder(transactionUsecase, itemUC, cartUsecase, db)

	}
	log.Print(app.Listen(port))
}

package http

import (
	"fmt"
	"os"
	"pos/internal/helper"
	itemHandler "pos/internal/item/delivery"
	itemrepository "pos/internal/item/repository"
	itemUsecase "pos/internal/item/usecase"
	"time"

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
	"github.com/gofiber/fiber/v2/middleware/logger"

	logs "log"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var validate = validator.New()

func HttpRun(port string) {
	myValidator := &internalValidator.XValidator{
		Validator: validate,
	}
	newLogger := gormLogger.New(
		logs.New(os.Stdout, "\r\n", logs.LstdFlags), // io writer
		gormLogger.Config{
			SlowThreshold:             time.Second,     // Slow SQL threshold
			LogLevel:                  gormLogger.Info, // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,            // Don't include params in the SQL log
			Colorful:                  true,            // Disable color
		},
	)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			viper.GetString("database.host"),
			viper.GetString("database.user"),
			viper.GetString("database.pass"),
			viper.GetString("database.name"), viper.GetString("database.port")),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{Logger: newLogger})

	if err != nil {
		panic(err)
	}
	//create enum
	helper.CreateStatusEnum(db)
	helper.CreateDebtStatus(db)

	//migration
	db.AutoMigrate(&model.Item{}, &model.Price{}, &model.Transaction{}, &model.Cart{}, &model.Customer{}, &model.CustomerDebt{})
	//creating trigger function after migrations
	// helper.CheckCustomer(db)
	// helper.CheckCustomerCountAfterUpdate(db)
	// helper.UpdateTotalPrice(db)
	//creating table trigger after migrations
	// helper.UpdateTransactionTrigger(db)
	// helper.CartTrigger(db)
	// helper.TransactionTrigger(db)

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency} ${body} ${resBody}\n",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	priceRepo := priceRepository.NewPricePresistentRepository(db)
	priceUC := priceUsecase.NewPriceUsecase(priceRepo)
	priceHandler.NewPriceHandler(app, myValidator, priceUC)

	itemRepo := itemrepository.NewItemPresistenRepository(db)
	itemUC := itemUsecase.NewItemUsecase(itemRepo, priceUC)
	itemHandler.NewItemHandler(app, myValidator, itemUC)

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

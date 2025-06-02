package seeder

import (
	"encoding/json"
	"pos/internal/cart"
	"pos/internal/item"
	"pos/internal/model"

	"pos/internal/transaction"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func TransactionSeeder(transaction transaction.TransactionUsecase, item item.ItemUsecase, cart cart.CartUsecase, db *gorm.DB) {
	customer := CustomerName(db)

	NewTrx := model.Transaction{
		CustomerID:   customer.ID,
		CustomerName: customer.CustomerName,
		Status:       "pending",
	}
	//check if transaction is not empty

	trx, count, _ := transaction.GetTransactions(0, 10, "", "pending", 1)
	if count > 0 {
		log.Print(trx)

		//check if cart is not empty
		if len(trx[0].Cart) <= 0 {
			SeedCart(item, cart, trx[0])
			return
		}

		for _, tx := range trx {
			log.Print(count)

			b, _ := json.Marshal(tx.Cart)
			log.Info().Bytes("cart", b)
			tx.CartJson.Scan(b)
			err := transaction.UpdateTransaction(&tx)
			if err != nil {
				panic(err)
			}
		}
		//get all trx again
		trx, _, _ := transaction.GetTransactions(0, 10, "", "pending", 1)
		log.Info().Interface("trx", trx)
		return
	}
	//create transaction first with status pending
	err := transaction.Savetransaction(&NewTrx)
	if err != nil {
		panic(err)
	}
	//
	//create cart based on transactionID
	//get item from item usecase

	SeedCart(item, cart, NewTrx)

}

func SeedCart(item item.ItemUsecase, cart cart.CartUsecase, NewTrx model.Transaction) {
	items, _, _ := item.GetItems(1, 10, "")
	log.Info().Any("items", items).Msg("Items ")
	for _, itemR := range items {

		quantity := randRange(1, 10)
		index := uint(randRange(0, len(itemR.Price)))
		cart.SaveCart(&model.Cart{
			TransactionID: NewTrx.ID,
			ItemID:        itemR.ID,
			Quantity:      quantity,
			ItemName:      itemR.ItemName,
			PriceID:       itemR.Price[index].ID,
			ItemPrice:     itemR.Price[index].Price,
			SubPrice:      itemR.Price[index].Price * int64(quantity),
			Unit:          itemR.Price[index].Unit,
		})
	}
}

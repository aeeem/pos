package seeder

import (
	"encoding/json"
	"pos/internal/cart"
	"pos/internal/item"
	"pos/internal/model"

	"pos/internal/transaction"

	"github.com/bxcodec/faker/v4"
	"github.com/rs/zerolog/log"
)

func TransactionSeeder(transaction transaction.TransactionUsecase, item item.ItemUsecase, cart cart.CartUsecase) {
	NewTrx := model.Transaction{
		CustomerName: faker.Name(),
		Status:       "pending",
	}
	//check if transaction is not empty

	trx, count, _ := transaction.GetTransactions(0, 10, "", "pending")
	log.Print(count)
	if count > 0 {
		log.Print(trx)

		//check if cart is not empty
		for _, tx := range trx {
			log.Print(count)

			b, _ := json.Marshal(tx.Cart)
			log.Info().Bytes("cart", b)
			tx.CartJson = &b
			err := transaction.UpdateTransaction(&tx)
			if err != nil {
				panic(err)
			}

		}
		//get all trx again
		trx, _, _ := transaction.GetTransactions(0, 10, "", "pending")
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
	items, _, _ := item.GetItems(1, 10, "")

	for _, itemR := range items {

		cart.Savetransaction(&model.Cart{
			TransactionID: NewTrx.ID,
			ItemID:        itemR.ID,
			Quantity:      1,
			ItemName:      itemR.ItemName,
			PriceID:       itemR.Price[uint(randRange(1, len(itemR.Price)))].ID,
		})
	}

}

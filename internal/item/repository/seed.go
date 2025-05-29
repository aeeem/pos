package repository

import (
	"math/rand"
	"pos/internal/item"

	"github.com/bxcodec/faker/v4"
	"gorm.io/gorm"
)

func SeedItem(db *gorm.DB) {
	for i := 0; i < 10; i++ {
		ItemModels := item.Item{
			ItemName:     faker.Word(),
			MaxPriceItem: int64(randRange(1, 5)),
		}
		err := db.Create(&ItemModels).Error
		if err != nil {
			panic(err)
		}
	}
}

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}

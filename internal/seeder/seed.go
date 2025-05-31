package seeder

import (
	"math/rand"
	"pos/internal/model"

	"github.com/bxcodec/faker/v4"
	"gorm.io/gorm"
)

func SeedItem(db *gorm.DB) {
	ic := int64(0)
	if err := db.Model(&model.Item{}).Count(&ic).Error; err != nil {
		panic(err)
	}

	if ic <= 0 {
		for i := 0; i < 10; i++ {
			ItemModels := model.Item{
				ItemName:     faker.Word(),
				MaxPriceItem: int64(randRange(1, 5)),
			}
			err := db.Create(&ItemModels).Error
			if err != nil {
				panic(err)
			}
		}
	} else {
		items := []model.Item{}
		db.Find(&items)
		for _, item := range items {
			Seedprice(db, int64(item.ID))
		}
	}

}
func Seedprice(db *gorm.DB, itemID int64) {
	ic := int64(0)
	if err := db.Model(&model.Price{}).Where("item_id = ?", itemID).Count(&ic).Error; err != nil {
		panic(err)
	}

	if ic <= 0 {
		for i := 0; i < 2; i++ {

			PriceModels := model.Price{
				Price:  int64(randRange(1000, 10000)),
				ItemID: uint(itemID),
				Active: true,
			}
			err := db.Create(&PriceModels).Error
			if err != nil {
				panic(err)
			}
		}

	}
	return
}

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}

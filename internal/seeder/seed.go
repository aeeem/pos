package seeder

import (
	"fmt"
	"math/rand"
	"pos/internal/model"
	"reflect"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func SeedItem(db *gorm.DB) {

	ic := int64(0)
	if err := db.Model(&model.Item{}).Count(&ic).Error; err != nil {
		panic(err)
	}
	faker := gofakeit.New(uint64(time.Now().Unix()))
	if ic <= 0 {
		for i := 0; i < 10; i++ {
			ItemModels := model.Item{
				ItemName:     faker.Fruit(),
				MaxPriceItem: int64(randRange(1000, 10000)),
			}

			err := db.Create(&ItemModels).Error
			if err != nil {
				if err == gorm.ErrDuplicatedKey {
					log.Print("duplicate key")
				}
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
		units := []string{"kg", "grams", "pcs"}
		for i := 0; i < 2; i++ {
			unit := randArrStr(units)
			PriceModels := model.Price{
				Price:  int64(randRange(1000, 10000)),
				ItemID: uint(itemID),
				Active: true,
				Unit:   reflect.ValueOf(unit).String(),
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
	log.Info().Msg(fmt.Sprint(min, " ", max))
	return rand.Intn(max-min) + min
}

func randArrStr(arr []string) string {

	min := 0
	max := len(arr)
	arrIndex := randRange(min, max)

	return arr[arrIndex]
}

func CustomerName(db *gorm.DB) model.Customer {
	Total := int64(0)
	cust := model.Customer{}
	err := db.Model(model.Customer{}).Count(&Total).Error

	err = db.Model(model.Customer{}).First(&cust).Error

	if Total > 0 {
		return cust
	}
	cust.CustomerName = "arif"
	cust.PhoneNumber = "08123456789"
	err = db.Create(&cust).Error
	if err != nil {
		panic(err)
	}
	return cust
}

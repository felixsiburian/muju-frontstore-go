package categories

import (
	"log"
	"muju-frontstore-go/database"
	"muju-frontstore-go/domain/model"
)

func GetCategories() []model.ProductCategory{
	var cat []model.ProductCategory
	var err error
	db := database.ConnectionDB()

	err = db.Debug().Find(&cat).Error
	if err != nil {
		log.Fatal()
	}

	return cat
}

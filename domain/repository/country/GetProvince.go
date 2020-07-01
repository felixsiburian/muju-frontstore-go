package country

import (
	"muju-frontstore-go/database"
	"muju-frontstore-go/domain/model"
)

func GetProvince() []model.Province{
	db := database.ConnectionDB()
	var prov []model.Province

	db.Find(&prov)
	return prov
}
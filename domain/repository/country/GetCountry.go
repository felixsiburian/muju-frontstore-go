package country

import (
	"muju-frontstore-go/database"
	"muju-frontstore-go/domain/model"
)

func GetCountry() []model.Country {
	var country []model.Country
	db := database.ConnectionDB()

	db.Find(&country)

	defer db.Close()
	return country
}

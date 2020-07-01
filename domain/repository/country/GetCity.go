package country

import (
	"database/sql"
	"fmt"
	"log"
	"muju-frontstore-go/database"
	"muju-frontstore-go/domain/model"
)

func GetCity(prov_id *int) []model.City {
	db := database.ConnectionDB()
	var city []model.City
	var _ *sql.Rows
	var err error
	
	if prov_id != nil {
		fmt.Println("masuk ke if")
		_, err = db.Find(&city).Where("province_id = 11").Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	db.Find(&city).Where("province_id = ?", 11)
	defer db.Close()
	return city
}
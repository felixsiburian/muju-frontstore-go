package main

import (
	"fmt"
	"muju-frontstore-go/database"
	"muju-frontstore-go/domain/model"
	"muju-frontstore-go/router"
)

func main() {
	fmt.Println("Welcome to Webserver")
	db := database.ConnectionDB()
	//Migrate table
	db.AutoMigrate(&model.Store{}, &model.Country{}, &model.PackageType{},
		&model.ProductCategory{}, &model.Template{}, &model.Province{}, &model.City{})
	e := router.New()
	e.Start(":8000")
}

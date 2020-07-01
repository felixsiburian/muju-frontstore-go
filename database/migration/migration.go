package main

import (
	"muju-frontstore-go/database"
	"muju-frontstore-go/domain/model"
)

func main()  {
	db := database.ConnectionDB()
	db.AutoMigrate(&model.Store{})
}
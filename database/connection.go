package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

func ConnectionDB() *gorm.DB {
	godotenv.Load(".env")
	db, err := gorm.Open(os.Getenv("DB_DRIVER"),os.Getenv("CONN"))
	if err != nil {
		fmt.Println("Failed Connect to Database", err.Error())
		panic("failed to connect to database")
	}

	return db
}

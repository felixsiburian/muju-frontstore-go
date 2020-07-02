package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/joho/godotenv"
	"os"
)

func ConnectionDB() *gorm.DB {
	godotenv.Load(".env")
	db, err := gorm.Open(os.Getenv("DB_DRIVER_SQL"),os.Getenv("CONN_SQL"))
	if err != nil {
		fmt.Println("Failed Connect to Database", err.Error())
		panic("failed to connect to database")
	}
	fmt.Println("Connected to Database : ", os.Getenv("DB_DRIVER_SQL"))
	return db
}

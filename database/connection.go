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
	db, err := gorm.Open("mssql", "sqlserver://moonlay:Standar123.@mulaijualan.database.windows.net?database=Muju-Frontstore")
	if err != nil {
		fmt.Println("Failed Connect to Database", err.Error())
		panic("failed to connect to database")
	}
	fmt.Println("Connected to Database : ", os.Getenv("DB_DRIVER_SQL"))
	return db
}

func ConnDB() *gorm.DB{
	db, err := gorm.Open("postgres","host=localhost port=5432 user=Muju dbname=Muju sslmode=disable password=Standar123.")
	if err != nil {
		fmt.Println("Failed Connect to Database", err.Error())
		panic("failed to connect to database")
	}

	return db
}

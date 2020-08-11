package main

import (
	"fmt"
	"github.com/beevik/guid"
	"muju-frontstore-go/database"
	"muju-frontstore-go/domain/model"
	"muju-frontstore-go/router"
	"time"
)

func main() {
	fmt.Println("Welcome to Webserver")

	/** your database connection here**/
	db := database.ConnectionDB()
	db.AutoMigrate(&model.Transaction{})
	db.Create(&model.Transaction{
		Id:            guid.NewString(),
		OrderId:       1001,
		StoreId:       1,
		StoreName:     "tmp",
		PackageId:     1,
		PackageName:   "tmp",
		TemplateId:    1,
		TemplateName:  "tmp",
		PackagePrice:  0,
		TemplatePrice: 0,
		InvoiceDate:   "2020-08-08",
		StartPeriod:   "2020-08-08",
		EndPeriod:     "2020-08-08",
		TotalPrice:    0,
		PaymentDate:   "2020-08-08",
		PaymentMethod: "OVO",
		CreatedBy:     "admin",
		CreatedDate:   time.Now(),
		ModifiedBy:    "admin",
		ModifiedDate:  time.Now(),
		DeletedBy:     "admin",
		DeletedDate:   time.Now(),
		Active:        true,
		IsDeleted:     false,
	})

	//Migrate table
	/** your migration code here**/

	e := router.New()
	e.Start(":8000")
}

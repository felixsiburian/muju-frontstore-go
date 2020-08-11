package transaction

import (
	"database/sql"
	"fmt"
	"github.com/beevik/guid"
	"log"
	"muju-frontstore-go/database"
	"muju-frontstore-go/domain/model"
	"time"
)

func CreateTransaction(transaction *model.Transaction) error {
	db := database.ConnectionDB()
	trsct := model.Transaction{
		Id:            guid.NewString(),
		StoreId:       transaction.StoreId,
		StoreName:     "tmp",
		PackageId:     transaction.PackageId,
		PackageName:   "tmp",
		TemplateId:    transaction.TemplateId,
		TemplateName:  "tmp",
		PackagePrice:  0,
		TemplatePrice: 0,
		InvoiceDate:   transaction.InvoiceDate,
		StartPeriod:   transaction.StartPeriod,
		EndPeriod:     transaction.EndPeriod,
		TotalPrice:    0,
		PaymentDate:   transaction.PaymentDate,
		PaymentMethod: transaction.PaymentMethod,
		CreatedBy:     "admin",
		CreatedDate:   time.Now(),
		ModifiedBy:    "admin",
		ModifiedDate:  time.Now(),
		DeletedBy:     "admin",
		DeletedDate:   time.Now(),
		Active:        true,
		IsDeleted:     false,
	}
	err := db.Debug().Create(&trsct).Error
	if err != nil {
		log.Fatal("Error Insert transaction")
	}

	defer db.Close()
	return err
}

func UpdateTransaction(transaction *model.Transaction) error {
	db := database.ConnectionDB()
	fmt.Println("orderId = ", transaction.OrderId)
	//os.Exit(1)
	err := db.Debug().Model(&transaction).Where("order_id = ?", transaction.OrderId ).Updates(map[string]interface{}{
		"store_id": transaction.StoreId,
		"package_id": transaction.PackageId,
		"template_id": transaction.TemplateId,
		"start_period": transaction.StartPeriod,
		"end_period": transaction.EndPeriod,
		"payment_method": transaction.PaymentMethod,
		"modified_date": time.Now(),
	}).Error
	if err != nil {
		log.Fatal("error update", err.Error())
	}

	defer db.Close()
	return err
}

func DeleteTransaction(transaction *model.Transaction) error {
	db := database.ConnectionDB()
	err := db.Debug().Model(&transaction).Where("order_id = ?", transaction.OrderId).Updates(map[string]interface{}{
		"active": false,
		"is_deleted": true,
		"deleted_date": time.Now(),
	}).Error
	if err != nil {
		log.Fatal("error delete : ", err.Error)
	}

	return err
}

func GetTransaction(page *int, size *int, store_id *int) []model.Transaction {
	db := database.ConnectionDB()
	var transactions []model.Transaction
	var rows *sql.Rows
	var total int
	var err error

	if page == nil && size == nil && store_id == nil {
		fmt.Println("masuk 1")
		rows, err = db.Find(&transactions).Order("created_date desc").Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil && store_id != nil {
		fmt.Println("masuk 2")
		rows, err = db.Find(&transactions).Where("store_id = ?", store_id).Count(total).Limit(*size).Offset(*page).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	result := make([]model.Transaction, 0)
	for rows.Next() {
		t := &model.Transaction{}
		err = rows.Scan(
			&t.Id,
			&t.OrderId,
			&t.StoreId,
			&t.StoreName,
			&t.PackageId,
			&t.PackageName,
			&t.TemplateId,
			&t.TemplateName,
			&t.PackagePrice,
			&t.TemplatePrice,
			&t.InvoiceDate,
			&t.StartPeriod,
			&t.EndPeriod,
			&t.TotalPrice,
			&t.PaymentDate,
			&t.PaymentMethod,
			&t.CreatedBy,
			&t.CreatedDate,
			&t.ModifiedBy,
			&t.ModifiedDate,
			&t.DeletedBy,
			&t.DeletedDate,
			&t.Active,
			&t.IsDeleted,
			)
		//store package template
		//package price template price & total price
		store := new(model.Store)
		db.Table("stores").Select("stores.store_name").
			Where("id = ?", t.StoreId).First(&store)
		t.StoreName = store.StoreName

		pkg := new(model.PackageType)
		db.Table("package_types").Select("package_types.package_name").
			Where("id = ?",t.PackageId).First(&pkg)
		t.PackageName = pkg.PackageName
		db.Table("package_types").Select("package_types.package_price").
			Where("id = ?",t.PackageId).First(&pkg)
		t.PackagePrice = pkg.PackagePrice

		tmp := new(model.Template)
		db.Table("templates").Select("templates.template_name").
			Where("id = ?", t.TemplateId).First(&tmp)
		t.TemplateName = tmp.TemplateName
		db.Table("templates").Select("templates.template_price").
			Where("id = ?", t.TemplateId).First(&tmp)
		t.TemplatePrice = tmp.TemplatePrice

		t.TotalPrice = t.PackagePrice + t.TemplatePrice

		result = append(result, *t)
	}
	defer db.Close()
	return result
}
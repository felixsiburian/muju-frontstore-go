package repository

import (
	"log"
	"muju-frontstore-go/database"
	"muju-frontstore-go/domain/model"
	"time"
)

func CreateStores(store *model.Store) error {
	db := database.ConnectionDB()
	stores := model.Store{
		StoreName:         store.StoreName,
		StoreDomain:       store.StoreDomain,
		ProductCategoryId: store.ProductCategoryId,
		CountryId:         store.CountryId,
		ProvinceId:        store.ProvinceId,
		CityId:            store.CityId,
		PostalCode:        store.PostalCode,
		CreatedBy:         "Admin",
		CreatedDate:       time.Now(),
		ModifiedBy:        "Admin",
		ModifiedDate:      time.Now(),
		DeletedBy:         "Admin",
		DeletedDate:       time.Now(),
		Active:            true,
		IsDeleted:         false,
	}
	err := db.Debug().Create(&stores).Error
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func UpdateStores(store *model.Store) error {
	db := database.ConnectionDB()
	err := db.Debug().Model(&store).Where("id = ?", store.Id).Updates(map[string]interface{}{
		"store_name":          store.StoreName,
		"store_domain":        store.StoreDomain,
		"product_category_id": store.ProductCategoryId,
		"country_id":          store.CountryId,
		"province_id":         store.ProvinceId,
		"city_id":             store.CityId,
		"postal_code":         store.PostalCode,
	}).Error
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func DeleteStores(store *model.Store) error {
	db := database.ConnectionDB()
	err := db.Debug().Model(&store).Where("id = ?", store.Id).Updates(map[string]interface{}{
		"active":     false,
		"is_deleted": true,
	}).Error
	if err != nil {
		log.Fatal(err)
	}
	store.DeletedDate = time.Now()
	return err
}

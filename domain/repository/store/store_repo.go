package store

import (
	"database/sql"
	"fmt"
	"log"
	"muju-frontstore-go/database"
	"muju-frontstore-go/domain/model"
	"time"
)

//type StoreRepository interface {
//	CreateStores(store *model.Store) error
//	UpdateStores(store *model.Store) error
//	DeleteStores(store *model.Store) error
//}

//type store_repo struct {
//	DB *gorm.DB
//}

//func Store_Repository(db *gorm.DB) StoreRepository {
//	return &store_repo{
//		DB: db,
//	}
//}

//func (s store_repo) CreateStores(store *model.Store) error {
//	db := database.ConnDB()
//	db.AutoMigrate(&model.Store{})
//	stores := model.Store{
//		StoreName:         store.StoreName,
//		StoreDomain:       store.StoreDomain,
//		ProductCategoryId: store.ProductCategoryId,
//		CountryId:         store.CountryId,
//		CountryName:       "tmp",
//		ProvinceId:        store.ProvinceId,
//		ProvinceName:      "tmp",
//		CityId:            store.CityId,
//		CityName:          "tmp",
//		PostalCode:        store.PostalCode,
//		CreatedBy:         "Admin",
//		CreatedDate:       time.Now(),
//		ModifiedBy:        "Admin",
//		ModifiedDate:      time.Now(),
//		DeletedBy:         "Admin",
//		DeletedDate:       time.Now(),
//		Active:            true,
//		IsDeleted:         false,
//	}
//	err := db.Debug().Create(&stores).Error
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//	return err
//}

func CreateStores(store *model.Store) error {
	db := database.ConnectionDB()
	stores := model.Store{
		StoreName:         store.StoreName,
		StoreDomain:       store.StoreDomain,
		ProductCategoryId: store.ProductCategoryId,
		CountryId:         store.CountryId,
		CountryName:       "tmp",
		ProvinceId:        store.ProvinceId,
		ProvinceName:      "tmp",
		CityId:            store.CityId,
		CityName:          "tmp",
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
	defer db.Close()
	return err
}

//func (s store_repo) UpdateStores(store *model.Store) error {
//	db := database.ConnDB()
//	err := db.Debug().Model(&store).Where("id = ?", store.Id).Updates(map[string]interface{}{
//		"store_name":          store.StoreName,
//		"store_domain":        store.StoreDomain,
//		"product_category_id": store.ProductCategoryId,
//		"country_id":          store.CountryId,
//		"province_id":         store.ProvinceId,
//		"city_id":             store.CityId,
//		"postal_code":         store.PostalCode,
//	}).Error
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//	return err
//}

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
	defer db.Close()
	return err
}

//func (s store_repo) DeleteStores(store *model.Store) error {
//	db := database.ConnDB()
//	err := db.Debug().Model(&store).Where("id = ?", store.Id).Updates(map[string]interface{}{
//		"active":     false,
//		"is_deleted": true,
//	}).Error
//	if err != nil {
//		log.Fatal(err)
//	}
//	store.DeletedDate = time.Now()
//	defer db.Close()
//	return err
//}

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
	defer db.Close()
	return err
}

func GetStore(page *int, size *int) []model.Store {
	db := database.ConnectionDB()
	var stores []model.Store
	var rows *sql.Rows
	var total int
	var err error

	if page == nil && size == nil {
		fmt.Println("masuk 1")
		rows, err = db.Find(&stores).Order("created_date desc").Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil {
		rows, err = db.Debug().Find(&stores).Order("created_date desc").Count(total).Limit(*size).Offset(*page).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	result := make([]model.Store, 0)
	for rows.Next() {
		s := &model.Store{}
		err = rows.Scan(
			&s.Id,
			&s.StoreName,
			&s.StoreDomain,
			&s.ProductCategoryId,
			&s.CountryId,
			&s.CountryName,
			&s.ProvinceId,
			&s.ProvinceName,
			&s.CityId,
			&s.CityName,
			&s.PostalCode,
			&s.CreatedBy,
			&s.CreatedDate,
			&s.ModifiedBy,
			&s.ModifiedDate,
			&s.DeletedBy,
			&s.DeletedDate,
			&s.Active,
			&s.IsDeleted,
		)

		country := new(model.Country)
		db.Table("countries").Select("countries.country").
			Where("id = ?", s.CountryId).First(&country)
		s.CountryName = country.Country

		province := new(model.Province)
		db.Table("provinces").Select("provinces.province_name").
			Where("id = ?", s.ProvinceId).First(&province)
		s.ProvinceName = province.ProvinceName

		city := new(model.City)
		db.Table("cities").Select("cities.city_name").
			Where("id = ?", s.CityId).First(&city)
		s.CityName = city.CityName

		result = append(result, *s)
	}
	defer db.Close()
	return result
}

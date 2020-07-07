package packageType

import (
	"database/sql"
	"log"
	"muju-frontstore-go/database"
	"muju-frontstore-go/domain/model"
	"time"
)

//type PackageRepository interface {
//	CreatePackage(packages *model.PackageType) error
//	UpdatePackage(packages *model.PackageType) error
//	DeletePackage(packages *model.PackageType) error
//}

//type pkg_repo struct {
//	DB *gorm.DB
//}

//func Pkg_Repository(db *gorm.DB) PackageRepository {
//	return &pkg_repo{
//		DB: db,
//	}
//}

func CreatePackage(packages *model.PackageType) error {
	db := database.ConnectionDB()
	pkg := model.PackageType{
		PackageName:  packages.PackageName,
		PackagePrice: packages.PackagePrice,
		CreatedDate:  time.Now(),
		CreatedBy:    "Admin",
		ModifiedDate: time.Now(),
		ModifiedBy:   "Admin",
		DeletedDate:  time.Now(),
		DeletedBy:    "Admin",
		Active:       true,
		IsDeleted:    false,
	}
	err := db.Debug().Create(&pkg).Error
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return err
}

//func (p pkg_repo) CreatePackage(packages *model.PackageType) error {
//	db := database.ConnDB()
//	db.AutoMigrate(&model.PackageType{})
//	pkg := model.PackageType{
//		PackageName:  packages.PackageName,
//		PackagePrice: packages.PackagePrice,
//		CreatedDate:  time.Now(),
//		CreatedBy:    "Admin",
//		ModifiedDate: time.Now(),
//		ModifiedBy:   "Admin",
//		DeletedDate:  time.Now(),
//		DeletedBy:    "Admin",
//		Active:       true,
//		IsDeleted:    false,
//	}
//	err := db.Debug().Create(&pkg).Error
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//	return err
//}

func UpdatePackage(packages *model.PackageType) error {
	db := database.ConnectionDB()
	err := db.Debug().Model(&packages).Updates(map[string]interface{}{
		"package_name":  packages.PackageName,
		"package_price": packages.PackagePrice,
		"modified_date": time.Now(),
	}).Error
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return err
}

//func (p pkg_repo) UpdatePackage(packages *model.PackageType) error {
//	db := database.ConnDB()
//	err := db.Debug().Model(&packages).Updates(map[string]interface{}{
//		"package_name":  packages.PackageName,
//		"package_price": packages.PackagePrice,
//		"modified_date": time.Now(),
//	}).Error
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//	return err
//}

func DeletePackage(packages *model.PackageType) error {
	db := database.ConnectionDB()
	err := db.Debug().Model(&packages).Updates(map[string]interface{}{
		"active":       false,
		"is_deleted":   true,
		"deleted_date": time.Now(),
	}).Error
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return err
}

//func (p pkg_repo) DeletePackage(packages *model.PackageType) error {
//	db := database.ConnDB()
//	err := db.Debug().Model(&packages).Updates(map[string]interface{}{
//		"active":       false,
//		"is_deleted":   true,
//		"deleted_date": time.Now(),
//	}).Error
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//	return err
//}

func GetPackage(page *int, size *int) []model.PackageType {
	db := database.ConnectionDB()
	var pkg []model.PackageType
	var rows *sql.Rows
	var total int
	var err error

	if page == nil && size == nil {
		rows, err = db.Find(&pkg).Order("created_date desc").Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil {
		rows, err = db.Debug().Find(&pkg).Order("created_date desc").Count(total).Limit(*size).Offset(*page).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	result := make([]model.PackageType, 0)
	for rows.Next() {
		p := &model.PackageType{}
		err = rows.Scan(
			&p.Id,
			&p.PackageName,
			&p.PackagePrice,
			&p.CreatedDate,
			&p.CreatedBy,
			&p.ModifiedDate,
			&p.ModifiedBy,
			&p.DeletedDate,
			&p.DeletedBy,
			&p.Active,
			&p.IsDeleted,
			)
		result = append(result, *p)
	}
	defer db.Close()
	return result
}

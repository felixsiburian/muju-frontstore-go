package Template

import (
	"database/sql"
	"log"
	"muju-frontstore-go/database"
	"muju-frontstore-go/domain/model"
	"time"
)

func CreateTemplate(templates *model.Template) error {
	db := database.ConnectionDB()
	template := model.Template{
		TemplateName:      templates.TemplateName,
		TemplatePrice:     templates.TemplatePrice,
		UrlDemo:           templates.UrlDemo,
		ProductCategoryId: templates.ProductCategoryId,
		ProductCategory:   "tmp",
		CreatedDate:       time.Now(),
		CreatedBy:         "Admin",
		ModifiedDate:      time.Now(),
		ModifiedBy:        "Admin",
		DeletedDate:       time.Now(),
		DeletedBy:         "Admin",
		Active:            true,
		IsDeleted:         false,
	}
	err := db.Debug().Create(&template).Error
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func UpdateTemplate(templates *model.Template) error {
	db := database.ConnectionDB()

	err := db.Debug().Model(&templates).Updates(map[string]interface{}{
		"template_name":  templates.TemplateName,
		"template_price": templates.TemplatePrice,
		"modified_date":  time.Now(),
	}).Error
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func DeleteTemplate(templates *model.Template) error {
	db := database.ConnectionDB()

	err := db.Debug().Model(&templates).Updates(map[string]interface{}{
		"active":       false,
		"is_deleted":   true,
		"deleted_date": time.Now(),
	}).Error

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func GetTemplate(page *int, size *int, cat *int) []model.Template {
	db := database.ConnectionDB()
	var temp []model.Template
	var rows *sql.Rows
	var total int
	var err error

	if page == nil && size == nil && cat == nil {
		rows, err = db.Debug().Find(&temp).Order("created_date desc").Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	if page != nil && size != nil && cat != nil {
		rows, err = db.Debug().Find(&temp).Where("product_category_id = ?", cat).Order("created_date desc").Count(total).Limit(*size).Offset(*page).Rows()
		if err != nil {
			log.Fatal(err)
		}
	}

	result := make([]model.Template, 0)
	for rows.Next() {
		t := &model.Template{}
		err = rows.Scan(
			&t.Id,
			&t.TemplateName,
			&t.TemplatePrice,
			&t.UrlDemo,
			&t.ProductCategoryId,
			&t.ProductCategory,
			&t.CreatedDate,
			&t.CreatedBy,
			&t.ModifiedDate,
			&t.ModifiedBy,
			&t.DeletedDate,
			&t.DeletedBy,
			&t.Active,
			&t.IsDeleted,
		)

		category := new(model.ProductCategory)
		db.Table("product_categories").Select("product_categories.product_category").
			Where("id = ?", t.ProductCategoryId).First(&category)
		t.ProductCategory = category.ProductCategory

		result = append(result, *t)
	}

	defer db.Close()
	return result
}

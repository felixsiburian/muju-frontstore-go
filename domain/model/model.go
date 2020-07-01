package model

import "time"

type Store struct {
	Id                int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	StoreName         string    `json:"store_name"`
	StoreDomain       string    `json:"store_domain"`
	ProductCategoryId int       `json:"product_category_id"`
	CountryId         int       `json:"country_id"`
	ProvinceId        int       `json:"province_id"`
	CityId            int       `json:"city_id"`
	PostalCode        string    `json:"postal_code"`
	CreatedBy         string    `json:"created_by"`
	CreatedDate       time.Time `json:"created_date"`
	ModifiedBy        string    `json:"modified_by"`
	ModifiedDate      time.Time `json:"modified_date"`
	DeletedBy         string    `json:"deleted_by"`
	DeletedDate       time.Time `json:"deleted_date"`
	Active            bool      `json:"active"`
	IsDeleted         bool      `json:"is_deleted"`
}

type Country struct {
	Id      int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Country string `json:"country"`
}

type Province struct {
	Id           int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	ProvinceName string `json:"province_name"`
}

type City struct {
	ProvinceId int    `json:"province_id"`
	Id         int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	CityName   string `json:"city_name"`
}

type Transaction struct {
	Id                int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	TransactionNumber int       `json:"transaction_number"`
	StoreId           int       `json:"store_id"`
	PackageId         int       `json:"package_id"`
	TemplateId        int       `json:"template_id"`
	PackagePrice      int       `json:"package_price"`
	TemplatePrice     int       `json:"template_price"`
	StartPeriod       string    `json:"start_period"`
	EndPeriod         string    `json:"end_period"`
	TotalPrice        int       `json:"total_price"`
	PaymentDate       string    `json:"payment_date"`
	PaymentMethod     string    `json:"payment_method"`
	CreatedBy         string    `json:"created_by"`
	CreatedDate       time.Time `json:"created_date"`
	ModifiedBy        string    `json:"modified_by"`
	ModifiedDate      time.Time `json:"modified_date"`
	DeletedBy         string    `json:"deleted_by"`
	DeletedAt         time.Time `json:"deleted_at"`
	Active            bool      `json:"active"`
	IsDeleted         bool      `json:"is_deleted"`
}

type PackageType struct {
	Id           int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	PackageName  string `json:"package_name"`
	PackagePrice int    `json:"package_price"`
}

type ProductCategory struct {
	Id              int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	ProductCategory string `json:"product_category"`
}

type Template struct {
	Id                int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	TemplateName      string `json:"template_name"`
	TemplatePrice     string `json:"template_price"`
	UrlDemo           string `json:"url_demo"`
	ProductCategoryId int    `json:"product_category_id"`
}

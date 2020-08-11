package model

import "time"

type Store struct {
	Id                int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	StoreName         string    `json:"store_name"`
	StoreDomain       string    `json:"store_domain"`
	ProductCategoryId int       `json:"product_category_id"`
	CountryId         int       `json:"country_id"`
	CountryName       string    `json:"country_name"`
	ProvinceId        int       `json:"province_id"`
	ProvinceName      string    `json:"province_name"`
	CityId            int       `json:"city_id"`
	CityName          string    `json:"city_name"`
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
	Id         int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL" json:"id"`
	CityName   string `json:"city_name"`
}

type Transaction struct {
	Id            string    `json:"id"`
	OrderId       int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL;" json:"order_id" `
	StoreId       int       `json:"store_id"`
	StoreName     string    `json:"store_name"`
	PackageId     int       `json:"package_id"`
	PackageName   string    `json:"package_name"`
	TemplateId    int       `json:"template_id"`
	TemplateName  string    `json:"template_name"`
	PackagePrice  int       `json:"package_price"`
	TemplatePrice int       `json:"template_price"`
	InvoiceDate   string    `json:"invoice_date"`
	StartPeriod   string    `json:"start_period"`
	EndPeriod     string    `json:"end_period"`
	TotalPrice    int       `json:"total_price"`
	PaymentDate   string    `json:"payment_date"`
	PaymentMethod string    `json:"payment_method"`
	CreatedBy     string    `json:"created_by"`
	CreatedDate   time.Time `json:"created_date"`
	ModifiedBy    string    `json:"modified_by"`
	ModifiedDate  time.Time `json:"modified_date"`
	DeletedBy     string    `json:"deleted_by"`
	DeletedDate     time.Time `json:"deleted_date"`
	Active        bool      `json:"active"`
	IsDeleted     bool      `json:"is_deleted"`
}

type PackageType struct {
	Id           int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	PackageName  string    `json:"package_name"`
	PackagePrice int       `json:"package_price"`
	CreatedDate  time.Time `json:"created_date"`
	CreatedBy    string    `json:"created_by"`
	ModifiedDate time.Time `json:"modified_date"`
	ModifiedBy   string    `json:"modified_by"`
	DeletedDate  time.Time `json:"deleted_date"`
	DeletedBy    string    `json:"deleted_by"`
	Active       bool      `json:"active"`
	IsDeleted    bool      `json:"is_deleted"`
}

type ProductCategory struct {
	Id              int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	ProductCategory string `json:"product_category"`
}

type Template struct {
	Id                int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	TemplateName      string    `json:"template_name"`
	TemplatePrice     int       `json:"template_price"`
	UrlDemo           string    `json:"url_demo"`
	ProductCategoryId int       `json:"product_category_id"`
	ProductCategory   string    `json:"product_category"`
	CreatedDate       time.Time `json:"created_date"`
	CreatedBy         string    `json:"created_by"`
	ModifiedDate      time.Time `json:"modified_date"`
	ModifiedBy        string    `json:"modified_by"`
	DeletedDate       time.Time `json:"deleted_date"`
	DeletedBy         string    `json:"deleted_by"`
	Active            bool      `json:"active"`
	IsDeleted         bool      `json:"is_deleted"`
}

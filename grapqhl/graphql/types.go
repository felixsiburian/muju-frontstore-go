package graphql

import "github.com/graphql-go/graphql"

var storeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Store",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"store_name": &graphql.Field{
			Type: graphql.String,
		},
		"store_domain": &graphql.Field{
			Type: graphql.String,
		},
		"product_category_id": &graphql.Field{
			Type: graphql.Int,
		},
		"country_id": &graphql.Field{
			Type: graphql.Int,
		},
		"country_name": &graphql.Field{
			Type: graphql.String,
		},
		"province_id": &graphql.Field{
			Type: graphql.Int,
		},
		"province_name": &graphql.Field{
			Type: graphql.String,
		},
		"city_id": &graphql.Field{
			Type: graphql.Int,
		},
		"city_name": &graphql.Field{
			Type: graphql.String,
		},
		"postal_code": &graphql.Field{
			Type: graphql.String,
		},
		"created_by": &graphql.Field{
			Type: graphql.String,
		},
		"created_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"modified_by": &graphql.Field{
			Type: graphql.String,
		},
		"modified_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type: graphql.String,
		},
		"deleted_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_deleted": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var cityType = graphql.NewObject(graphql.ObjectConfig{
	Name: "City",
	Fields: graphql.Fields{
		"province_id": &graphql.Field{
			Type: graphql.Int,
		},
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"city_name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var provinceType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Province",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"province_name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var countryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Country",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"country": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var packageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Package",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"package_name": &graphql.Field{
			Type: graphql.String,
		},
		"package_price": &graphql.Field{
			Type: graphql.Int,
		},
		"created_by": &graphql.Field{
			Type: graphql.String,
		},
		"created_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"modified_by": &graphql.Field{
			Type: graphql.String,
		},
		"modified_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type: graphql.String,
		},
		"deleted_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_deleted": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var categoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Category",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"product_category": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var templateType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Template",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"template_name": &graphql.Field{
			Type: graphql.String,
		},
		"template_price": &graphql.Field{
			Type: graphql.Int,
		},
		"url_demo": &graphql.Field{
			Type: graphql.String,
		},
		"product_category_id": &graphql.Field{
			Type: graphql.Int,
		},
		"product_category": &graphql.Field{
			Type: graphql.String,
		},
		"created_by": &graphql.Field{
			Type: graphql.String,
		},
		"created_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"modified_by": &graphql.Field{
			Type: graphql.String,
		},
		"modified_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type: graphql.String,
		},
		"deleted_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_deleted": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var transactionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Transaction",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"order_id": &graphql.Field{
			Type: graphql.Int,
		},
		"store_id": &graphql.Field{
			Type: graphql.Int,
		},
		"store_name": &graphql.Field{
			Type: graphql.String,
		},
		"package_id": &graphql.Field{
			Type: graphql.Int,
		},
		"package_name": &graphql.Field{
			Type: graphql.String,
		},
		"template_id": &graphql.Field{
			Type: graphql.Int,
		},
		"template_name": &graphql.Field{
			Type: graphql.String,
		},
		"package_price": &graphql.Field{
			Type: graphql.Int,
		},
		"template_price": &graphql.Field{
			Type: graphql.Int,
		},
		"invoice_date": &graphql.Field{
			Type: graphql.String,
		},
		"start_period": &graphql.Field{
			Type: graphql.String,
		},
		"end_period": &graphql.Field{
			Type: graphql.String,
		},
		"total_price": &graphql.Field{
			Type: graphql.Int,
		},
		"payment_date": &graphql.Field{
			Type: graphql.String,
		},
		"payment_method": &graphql.Field{
			Type: graphql.String,
		},
		"created_by": &graphql.Field{
			Type: graphql.String,
		},
		"created_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"modified_by": &graphql.Field{
			Type: graphql.String,
		},
		"modified_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"deleted_by": &graphql.Field{
			Type: graphql.String,
		},
		"deleted_date": &graphql.Field{
			Type: graphql.DateTime,
		},
		"active": &graphql.Field{
			Type: graphql.Boolean,
		},
		"is_deleted": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})
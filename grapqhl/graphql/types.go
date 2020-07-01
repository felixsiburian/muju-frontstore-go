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
		"province_id": &graphql.Field{
			Type: graphql.Int,
		},
		"city_id": &graphql.Field{
			Type: graphql.Int,
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
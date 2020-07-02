package graphql

import "github.com/graphql-go/graphql"

type Root struct {
	Query *graphql.Object
}

type Page struct {
	page int `json:"null"`
}

func NewRoot() *Root {
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"store": &graphql.Field{
						Type: graphql.NewList(storeType),
						Args: graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: StoreResolver,
					},
					"city": &graphql.Field{
						Type: graphql.NewList(cityType),
						Args: graphql.FieldConfigArgument{
							"prov_id": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: CityResolver,
					},
					"province": &graphql.Field{
						Type:    graphql.NewList(provinceType),
						Resolve: ProvinceResolver,
					},
					"country": &graphql.Field{
						Type:    graphql.NewList(countryType),
						Resolve: CountryResolver,
					},
					"package": &graphql.Field{
						Type: graphql.NewList(packageType),
						Args: graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: PackageResolve,
					},
					"template": &graphql.Field{
						Type: graphql.NewList(templateType),
						Args: graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"cat": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: TemplateResolver,
					},
					"category": &graphql.Field{
						Type: graphql.NewList(categoryType),
						Args: graphql.FieldConfigArgument{
						},
						Resolve: CategoryResolver,
					},
				},
			},
		),
	}
	return &root
}

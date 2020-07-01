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
				},
			},
		),
	}
	return &root
}

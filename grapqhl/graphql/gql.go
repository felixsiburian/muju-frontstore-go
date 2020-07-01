package graphql

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

func ExecuteQuery(query string) *graphql.Result{

	rootQuery := NewRoot()
	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)
	result := graphql.Do(graphql.Params{
		Schema:schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("Unexpected errors inside Execute Query : %v", result.Errors)
	}
	return result
}
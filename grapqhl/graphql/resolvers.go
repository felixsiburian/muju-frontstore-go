package graphql

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"muju-frontstore-go/domain/repository"
)

func StoreResolver(p graphql.ResolveParams)(interface{}, error) {
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	if ok && sip {
		var pages *int = &page
		var sizes *int = &size
		stores := repository.GetStore(pages, sizes)
		fmt.Println("stores : ", stores)
		return stores, nil
	}
	stores := repository.GetStore(nil, nil)
	return stores,nil
}
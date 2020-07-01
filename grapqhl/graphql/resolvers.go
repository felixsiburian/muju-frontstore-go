package graphql

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"muju-frontstore-go/domain/repository/store"
	"muju-frontstore-go/domain/repository/country"
)

func StoreResolver(p graphql.ResolveParams)(interface{}, error) {
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	if ok && sip {
		var pages *int = &page
		var sizes *int = &size
		stores := store.GetStore(pages, sizes)
		fmt.Println("stores : ", stores)
		return stores, nil
	}
	stores := store.GetStore(nil, nil)
	return stores,nil
}

func CityResolver(p graphql.ResolveParams)(interface{}, error) {
	prov_id, id := p.Args["prov_id"].(int)

	if id {
		var prov_ids *int = &prov_id
		city := country.GetCity(prov_ids)
		fmt.Println("City : ", city)
		return city, nil
 	}

	city := country.GetCity( nil)
	return city, nil
}

func ProvinceResolver(p graphql.ResolveParams)(interface{}, error){
	prov := country.GetProvince()
	return prov, nil
}

func CountryResolver(p graphql.ResolveParams)(interface{}, error){
	country := country.GetCountry()
	return country,nil
}
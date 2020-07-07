package graphql

import (
	"fmt"
	"github.com/graphql-go/graphql"
	tmp "muju-frontstore-go/domain/repository/Template"
	"muju-frontstore-go/domain/repository/categories"
	"muju-frontstore-go/domain/repository/country"
	pkg "muju-frontstore-go/domain/repository/packageType"
	"muju-frontstore-go/domain/repository/store"
)

func StoreResolver(p graphql.ResolveParams) (interface{}, error) {
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
	return stores, nil
}

func CityResolver(p graphql.ResolveParams) (interface{}, error) {
	prov_id, id := p.Args["prov_id"].(int)

	if id {
		var prov_ids *int = &prov_id
		city := country.GetCity(prov_ids)
		fmt.Println("City : ", city)
		return city, nil
	}

	city := country.GetCity(nil)
	return city, nil
}

func ProvinceResolver(p graphql.ResolveParams) (interface{}, error) {
	prov := country.GetProvince()
	return prov, nil
}

func CountryResolver(p graphql.ResolveParams) (interface{}, error) {
	country := country.GetCountry()
	return country, nil
}

func PackageResolve(p graphql.ResolveParams) (interface{}, error) {
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	if ok && sip {
		var pages *int = &page
		var sizes *int = &size
		pkg := pkg.GetPackage(pages, sizes)
		fmt.Println("Pkg : ", pkg)
		return pkg, nil
	}
	pkgs := pkg.GetPackage(nil, nil)
	return pkgs, nil
}

func TemplateResolver(p graphql.ResolveParams)(interface{}, error){
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	cat, cats := p.Args["cat"].(int)

	if ok && sip && cats {
		fmt.Println("masuk sini")
		var pages *int = &page
		var sizes *int = &size
		var catss *int = &cat
		tmp := tmp.GetTemplate(pages, sizes, catss)
		fmt.Println("Tmp : ", tmp)
		return tmp, nil
	}

	tmps := tmp.GetTemplate(nil, nil, nil)
	return tmps, nil
}

func CategoryResolver(p graphql.ResolveParams)(interface{}, error) {
	cat := categories.GetCategories()
	return cat, nil
}
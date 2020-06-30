package router

import (
	"github.com/labstack/echo"
	"muju-frontstore-go/kafka/API/Store"
	"muju-frontstore-go/kafka/Host"
)

func New() *echo.Echo {
	e := echo.New()

	Host.StartkafkaStore()

	e.POST("/CreateStores", Store.PublishCreateStore)
	e.POST("/DeleteStores", Store.PublishDeleteStore)
	e.POST("/UpdateStores", Store.PublishUpdateStore)
	return e
}
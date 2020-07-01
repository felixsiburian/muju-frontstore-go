package router

import (
	"crypto/subtle"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"muju-frontstore-go/kafka/API/Store"
	"muju-frontstore-go/kafka/Host"
	grapqhl "muju-frontstore-go/grapqhl"
)

func New() *echo.Echo {
	e := echo.New()

	Host.StartkafkaStore()

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("admin")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("Standar123.")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	grapqhl.MainGroup(e)

	e.POST("/CreateStores", Store.PublishCreateStore)
	e.POST("/DeleteStores", Store.PublishDeleteStore)
	e.POST("/UpdateStores", Store.PublishUpdateStore)
	return e
}
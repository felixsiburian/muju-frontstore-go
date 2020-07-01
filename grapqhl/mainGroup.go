package grapqhl

import (
	"github.com/labstack/echo"
	"muju-frontstore-go/handlers"
)

func MainGroup(e *echo.Echo) {
	e.POST("/graphql", handlers.Query)
}
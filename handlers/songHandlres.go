package handlers

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	grapqhl "muju-frontstore-go/grapqhl/graphql"
	"net/http"
)

type Server struct {
	GqlSchema *graphql.Schema
}

type ReqBody struct {
	Query string `json:"query"`
}

func Query(c echo.Context) error {
	var rBody ReqBody

	err := json.NewDecoder(c.Request().Body).Decode(&rBody)
	if err != nil {
		log.Printf("Failed Processing request: %s\n",err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	result := grapqhl.ExecuteQuery(rBody.Query)
	return c.JSON(http.StatusOK, result)
}
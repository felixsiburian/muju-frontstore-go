package main

import (
	"fmt"
	"muju-frontstore-go/router"
)

func main() {
	fmt.Println("Welcome to Webserver")
	e := router.New()
	e.Start(":8000")
}

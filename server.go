package main

import (
	"github.com/charlesfan/go-api/route"
	"github.com/charlesfan/go-api/service/rsi"
)

func main() {
	// rsi.Services init
	rsi.Init()
	router := route.Init()

	router.Run(":8080")
}

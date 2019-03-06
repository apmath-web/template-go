package main

import (
	"github.com/apmath-web/template-go/routing"
	"log"
)

func main() {
	router := routing.GenRouter()
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}

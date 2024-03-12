package main

import (
	"assignment-3/helpers"
	"assignment-3/routers"
)

func main() {
	go helpers.WriteJson()
	routers.Router()
}

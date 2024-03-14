package main

import (
	"assignment-3/helpers"
	"assignment-3/routers"
	"log"
	"os"
)

func init() {
	cwd, _ := os.Getwd()
	_, err := os.ReadFile(cwd + "\\test.json")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("test.json exist")
}

func main() {
	go helpers.WriteJson()
	routers.Router()
}

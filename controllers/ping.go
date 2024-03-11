package controllers

import (
	"assignment-3/structs"
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	var datas structs.Data
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	content, err := os.ReadFile(cwd + "\\test.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &datas)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"message": datas,
	})
}

package routers

import (
	"assignment-3/controllers"

	"github.com/gin-gonic/gin"
)

func Router() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.LoadHTMLFiles("index.tmpl")
	router.StaticFile("/meteocons--dust-wind-fill.svg", "./src/meteocons--dust-wind-fill.svg")
	router.Static("/assets/", "./src/")
	router.GET("/ping", controllers.Ping)
	router.GET("/", controllers.Index)
	router.Run(":8080")
}

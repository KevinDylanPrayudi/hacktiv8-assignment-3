<<<<<<< HEAD
package routers

import (
	"assignment-3/controllers"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.LoadHTMLFiles("index.tmpl")
	router.StaticFile("/meteocons--dust-wind-fill.svg", "./src/meteocons--dust-wind-fill.svg")
	router.Static("/assets/", "./src/")
	router.GET("/index", controllers.Index)
	router.GET("/ping", controllers.Ping)
	router.Run(":8080")
}
=======
package routers

import (
	"assignment-3/controllers"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.LoadHTMLFiles("index.tmpl")
	router.StaticFile("/meteocons--dust-wind-fill.svg", "./src/meteocons--dust-wind-fill.svg")
	router.Static("/assets/", "./src/")
	router.GET("/index", controllers.Index)
	router.GET("/ping", controllers.Ping)
	router.Run(":8080")
}
>>>>>>> 746ac3f6573f8d4f94917082e2a5d96c58cdf9f1

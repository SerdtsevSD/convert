package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

            
func main() {
	// redisFunc()
	router = gin.Default()
	router.Use(LiberalCORS)
	router.Delims("&{", "}&")
	router.LoadHTMLGlob("templates/*")
	router.StaticFS("/static", http.Dir("./static"))
	initializeRoutes()
	router.Run(":9090")

}

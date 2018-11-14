package main

import (
	"github.com/gin-gonic/gin"
)

func init() {
	//var err error
}
func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/", contorollers.helloHandler)
}

package main

import (
	"jcg-project/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
)

func init() {
	//var err error
}
func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/hello", controller.HelloHandler)

	http.Handle("/", router)
	appengine.Main()
}

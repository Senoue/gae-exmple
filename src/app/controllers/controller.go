package controllers

import (
	"log"

	"../models"
	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	var user models.User

	log.Println(user.ID)
}

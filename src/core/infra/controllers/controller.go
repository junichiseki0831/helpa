package controllers

import (
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
    r := gin.Default()

	r.GET("/health", getting)

	return r
}

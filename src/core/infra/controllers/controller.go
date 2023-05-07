package controllers

import (
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", getting)
	r.POST("/users", func(ctx *gin.Context) {
		newCreateUserController().createUserHandler(ctx)
	})

	return r
}

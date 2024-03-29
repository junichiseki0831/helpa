package controllers

import (
	"helpa/src/core/infra/middleware"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.HandleErrorMiddleware())

	r.GET("/health", getting)
	r.POST("/users", func(ctx *gin.Context) {
		newCreateUserController().createUserHandler(ctx)
	})

	return r
}

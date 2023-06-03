package controllers

import (
	"helpa/src/core/infra/middleware"

	"github.com/gin-gonic/gin"
)

func CreateRouter() (*gin.Engine, error) {
	r := gin.Default()

	db, err := middleware.NewDB()
	if err != nil {
		return nil, err
	}

	r.Use(middleware.HandleErrorMiddleware())
	r.Use(middleware.DBInterceptor(db))

	r.GET("/health", getting)
	r.POST("/users", func(ctx *gin.Context) {
		newCreateUserController().createUserHandler(ctx, db)
	})

	return r, nil
}

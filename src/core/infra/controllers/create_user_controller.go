package controllers

import (
	app "helpa/src/core/app/userapp"
	infra "helpa/src/core/infra/rdbimpl"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createUserController struct {
}

func newCreateUserController() *createUserController {
	return &createUserController{}
}

func (c *createUserController) createUserHandler(ctx *gin.Context) {
	var user app.CreateUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := infra.NewDB()
	if err != nil {
		ctx.Error(err)
		return
	}

	userRepoImpl := infra.NewUserRepositoryImpl(db)

	if err := app.NewCreateUserAppService(userRepoImpl).Exec(ctx, &user); err != nil {
		ctx.Error(err)
		return
	}
}

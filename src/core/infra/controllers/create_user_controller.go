package controllers

import (
	app "helpa/src/core/app/userapp"
	infra "helpa/src/core/infra/rdbimpl"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type createUserController struct {
}

func newCreateUserController() *createUserController {
	return &createUserController{}
}

func (c *createUserController) createUserHandler(ctx *gin.Context, db *sqlx.DB) {
	var user app.CreateUserRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(err)
		return
	}

	userRepoImpl := infra.NewUserRepositoryImpl(db)

	if err := app.NewCreateUserAppService(userRepoImpl).Exec(ctx, &user); err != nil {
		ctx.Error(err)
		return
	}
}

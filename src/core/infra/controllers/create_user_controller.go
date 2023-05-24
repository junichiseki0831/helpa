package controllers

import (
	app "helpa/src/core/app/userapp"
	infra "helpa/src/core/infra/rdbimpl"
	"helpa/src/support/smperr"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": smperr.BadRequest("Failed to bind JSON data")})
		return
	}

	db, err := infra.NewDB()
	if err != nil {
		ctx.Error(smperr.Internalf("Failed to initialize database: %w", err))
		return
	}

	userRepoImpl := infra.NewUserRepositoryImpl(db)

	if err := app.NewCreateUserAppService(userRepoImpl).Exec(ctx, &user); err != nil {
		smperr.Internalf("Failed to execute CreateUserAppService: %w", err)
		return
	}
}

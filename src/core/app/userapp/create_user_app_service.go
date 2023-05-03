package app

import (
	"context"
	"helpa/src/core/domain/shared/vo"
	domain "helpa/src/core/domain/userdm"
)

type CreateUserAppService struct {
	userRepo domain.UserRepository
}

func NewCreateUserAppService(userRepo domain.UserRepository) *CreateUserAppService {
	return &CreateUserAppService{
		userRepo,
	}
}

type CreateUserRequest struct {
	Name         string
	Password     vo.Password
	Email        vo.Email
	Introduction string
	Note         string
	Image        vo.Image
}

func (app *CreateUserAppService) Exec(ctx context.Context, req *CreateUserRequest) error {
	createUser, err := domain.GenWhenCreate(req.Name, req.Password, req.Email, req.Introduction, req.Note, req.Image)
	if err != nil {
		return err
	}
	return app.userRepo.Store(ctx, createUser)
}

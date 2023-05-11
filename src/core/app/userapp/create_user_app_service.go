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
	Password     string
	Email        string
	Introduction string
	Note         string
	Image        string
}

func (app *CreateUserAppService) Exec(ctx context.Context, req *CreateUserRequest) error {
	pass, err := vo.NewPassword(req.Password)
	if err != nil {
		return err
	}
	mail, err := vo.NewEmail(req.Email)
	if err != nil {
		return err
	}
	img, err := vo.NewImage(req.Image)
	if err != nil {
		return err
	}
	createUser, err := domain.GenWhenCreate(req.Name, pass, mail, req.Introduction, req.Note, img)
	if err != nil {
		return err
	}
	return app.userRepo.Store(ctx, createUser)
}

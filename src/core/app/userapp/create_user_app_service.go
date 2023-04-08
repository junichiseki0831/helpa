package app

import (
	"context"
	"encoding/base64"
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
	Image        []byte
}

func (app *CreateUserAppService) Exec(ctx context.Context, req *CreateUserRequest) error {
	base64String := base64.StdEncoding.EncodeToString(req.Image)
	createUser, err := domain.GenWhenCreate(req.Name, req.Password, req.Email, req.Introduction, req.Note, base64String)
	if err != nil {
		return err
	}
	return app.userRepo.Store(ctx, createUser)
}

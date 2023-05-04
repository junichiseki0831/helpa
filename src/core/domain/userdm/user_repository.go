package domain

import (
	"context"
)

type UserRepository interface {
	Store(ctx context.Context, user *User) error
	FindByName(ctx context.Context, name string) ([]User, error)
}

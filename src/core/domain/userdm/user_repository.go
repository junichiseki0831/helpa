package domain

import "context"

type UserRepository interface {
	Store(ctx context.Context) error
	FindByName(ctx context.Context, name string) (*User, error)
}

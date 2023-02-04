package domain

type UserRepository interface {
    Store() error
	FindByName() (*User, error)
}

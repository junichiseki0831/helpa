package infra

import (
	"context"
	_ "embed"
	domain "helpa/src/core/domain/userdm"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepositoryImpl(db *sqlx.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (repo *UserRepositoryImpl) FindByName(ctx context.Context, name string) ([]*domain.User, error) {
	db, err := NewDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var users []*domain.User
	err = db.Select(&users, "SELECT * FROM users WHERE name = ?", name)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *UserRepositoryImpl) Store(ctx context.Context, user *domain.User) error {
	db, err := NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(
		"INSERT INTO users (id, name, password, email, introduction, note, image, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		user.ID(), user.Name(), user.Password(), user.Email(), user.Introduction(), user.Note(), user.Image(), user.CreatedAt(), user.UpdatedAt(),
	)
	if err != nil {
		return err
	}

	return nil
}

func NewDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", "user:password@tcp(host:port)/dbname")
	if err != nil {
		return nil, err
	}
	return db, nil
}

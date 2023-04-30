package infra

import (
	"context"
	_ "embed"
	"helpa/src/core/domain/datamodel"
	domain "helpa/src/core/domain/userdm"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepositoryImpl(db *sqlx.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (repo *UserRepositoryImpl) FindByName(ctx context.Context, name string) ([]domain.User, error) {

	query := `SELECT *
	FROM users
	WHERE name = ?;`

	var userDMs []*datamodel.User
	if err := repo.db.Select(&userDMs, query, name); err != nil {
		return nil, err
	}

	users := make([]domain.User, 0, len(userDMs))
	for _, userDM := range userDMs {
		user, err := toDomainUser(userDM)
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}
	return users, nil
}

// func (repo *UserRepositoryImpl) Store(ctx context.Context, user *domain.User) error {
// 	db, err := NewDB()
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	_, err = db.Exec(
// 		"INSERT INTO users (id, name, password, email, introduction, note, image, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
// 		user.ID(), user.Name(), user.Password(), user.Email(), user.Introduction(), user.Note(), user.Image(), user.CreatedAt(), user.UpdatedAt(),
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func NewDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", "root:secrets@tcp(db:3306)/helpa")
	if err != nil {
		return nil, err
	}
	return db, nil
}

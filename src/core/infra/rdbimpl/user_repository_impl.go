package infra

import (
	"context"
	domain "helpa/src/core/domain/userdm"
	"helpa/src/core/infra/datamodel"
	"helpa/src/support/smperr"

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
		return nil, smperr.Internalf("Failed to query user data: %w", err)
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

func (repo *UserRepositoryImpl) Store(ctx context.Context, user *domain.User) error {
	query := `
        INSERT INTO users (id, name, password, email, introduction, note, image, createdAt, updatedAt)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

	_, err := repo.db.Exec(query,
		user.ID(),
		user.Name(),
		user.Password(),
		user.Email(),
		user.Introduction(),
		user.Note(),
		user.Image().Base64(),
		user.CreatedAt().String(),
		user.UpdatedAt().String(),
	)
	if err != nil {
		return smperr.Internalf("Failed to execute INSERT query: %w", err)
	}
	return nil
}

func NewDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", "root:secrets@tcp(db:3306)/helpa")
	if err != nil {
		return nil, smperr.Internalf("Failed to initialize database: %w", err)
	}
	return db, nil
}

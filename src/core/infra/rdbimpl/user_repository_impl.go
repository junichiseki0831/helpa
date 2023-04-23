package infra

import (
	"context"
	_ "embed"
	"helpa/src/core/domain/datamodel"
	"helpa/src/core/domain/shared/vo"
	domain "helpa/src/core/domain/userdm"
	"time"

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

	var users []domain.User
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

func toDomainUser(dmUser *datamodel.User) (*domain.User, error) {
	user := &domain.User{}
	user.SetID(domain.UserID(dmUser.ID))
	user.SetName(dmUser.Name)
	user.SetPassword(vo.Password(dmUser.Password))
	user.SetEmail(vo.Email(dmUser.Email))
	user.SetIntroduction(dmUser.Introduction)
	user.SetNote(dmUser.Note)
	if dmUser.Image.Valid {
		image, err := vo.NewImage(dmUser.Image.String)
		if err != nil {
			return nil, err
		}
		user.SetImage(image)
	}
	createdAt, err := time.Parse("2006-01-02 15:04:05", dmUser.CreatedAt)
	if err != nil {
		return nil, err
	}
	user.SetCreatedAt(vo.CreatedAt(createdAt))
	updatedAt, err := time.Parse("2006-01-02 15:04:05", dmUser.UpdatedAt)
	if err != nil {
		return nil, err
	}
	user.SetUpdatedAt(vo.UpdatedAt(updatedAt))

	return user, nil
}

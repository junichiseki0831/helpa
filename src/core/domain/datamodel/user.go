package datamodel

import (
	"database/sql"
)

type User struct {
	ID           string         `db:"id"`
	Name         string         `db:"name"`
	Password     string         `db:"password"`
	Email        string         `db:"email"`
	Introduction string         `db:"introduction"`
	Note         string         `db:"note"`
	Image        sql.NullString `db:"image"`
	CreatedAt    string         `db:"createdAt"`
	UpdatedAt    string         `db:"updatedAt"`
}

package domain

import (
	"helpa/src/core/domain/shared/vo"
)

type User struct {
	id           UserID
	name         string
	password     string
	email        string
	introduction string
	note         string
	externalLink string
	imege        string
	CreatedAt    vo.CreatedAt
	UpdatedAt    vo.UpdatedAt
}

func (u *User) ID() UserID {
	return u.id
}

package domain

import (
	"helpa/src/core/domain/shared/vo"
)

type User struct {
	id           UserID
	name         string
	password     vo.Password
	email        vo.Email
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

func (u *User) Password() vo.Password {
	return u.password
}

func (u *User) Email() vo.Email {
	return u.email
}

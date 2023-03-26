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
	createdAt    vo.CreatedAt
	updatedAt    vo.UpdatedAt
}

func newUser(
	id UserID,
	name string,
	password vo.Password,
	email vo.Email,
	introduction string,
	note string,
	externalLink string,
	imege string,
	createdAt vo.CreatedAt,
	updatedAt vo.UpdatedAt,
) *User {
	return &User{
		id:           id,
		name:         name,
		password:     password,
		email:        email,
		introduction: introduction,
		note:         note,
		externalLink: externalLink,
		imege:        imege,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Password() vo.Password {
	return u.password
}

func (u *User) Email() vo.Email {
	return u.email
}

func (u *User) Introduction() string {
	return u.introduction
}

func (u *User) Note() string {
	return u.note
}

func (u *User) ExternalLink() string {
	return u.externalLink
}

func (u *User) Imege() string {
	return u.imege
}

func (u *User) CreatedAt() vo.CreatedAt {
	return u.createdAt
}

func (u *User) UpdatedAt() vo.UpdatedAt {
	return u.updatedAt
}

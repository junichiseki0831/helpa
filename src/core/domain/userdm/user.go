package domain

import (
	"errors"
	"helpa/src/core/domain/shared/vo"
	"net/http"
	"unicode/utf8"
)

type User struct {
	id           UserID
	name         string
	password     vo.Password
	email        vo.Email
	introduction string
	note         string
	image        []byte
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
	image []byte,
	createdAt vo.CreatedAt,
	updatedAt vo.UpdatedAt,
) (*User, error) {

	nameLen := utf8.RuneCountInString(name)
	introductionLen := utf8.RuneCountInString(introduction)
	noteLen := utf8.RuneCountInString(note)
	if name == "" {
		return nil, errors.New("empty name")
	} else if nameLen > 50 {
		return nil, errors.New("please enter your name within 50 characters")
	}

	if introductionLen > 500 {
		return nil, errors.New("please enter your introduction within 500 characters")
	}

	if noteLen > 500 {
		return nil, errors.New("please enter your note within 500 characters")
	}

	mine := http.DetectContentType(image)
	if mine != "image/jpeg" && mine != "image/png" {
		return nil, errors.New("please specify the image extension as jpg or png")
	}

	return &User{
		id:           id,
		name:         name,
		password:     password,
		email:        email,
		introduction: introduction,
		note:         note,
		image:        image,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}, nil
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

func (u *User) Image() []byte {
	return u.image
}

func (u *User) CreatedAt() vo.CreatedAt {
	return u.createdAt
}

func (u *User) UpdatedAt() vo.UpdatedAt {
	return u.updatedAt
}

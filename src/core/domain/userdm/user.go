package domain

import (
	"errors"
	"helpa/src/core/domain/shared/vo"
	"net/http"
	"net/url"
)

type User struct {
	id           UserID
	name         string
	password     vo.Password
	email        vo.Email
	introduction string
	note         string
	externalLink string
	image        string
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
	image string,
	createdAt vo.CreatedAt,
	updatedAt vo.UpdatedAt,
) (*User, error) {

	nameLen := len(name)
	introductionLen := len(introduction)
	noteLen := len(note)
	if name == "" {
		return nil, errors.New("Empty name")
	} else if nameLen > 50 {
		return nil, errors.New("Please enter your name within 50 characters")
	}

	if introductionLen > 500 {
		return nil, errors.New("Please enter your introduction within 500 characters")
	}

	if noteLen > 500 {
		return nil, errors.New("Please enter your note within 500 characters")
	}

	_, err := url.ParseRequestURI(externalLink)
	if err != nil {
		return nil, errors.New("there is an error in the url")
	}

	mine := http.DetectContentType([]byte(image))
	if mine != "image/jpeg" || mine != "image/png" {
		return nil, errors.New("Please specify the image extension as jpg or png")
	}

	return &User{
		id:           id,
		name:         name,
		password:     password,
		email:        email,
		introduction: introduction,
		note:         note,
		externalLink: externalLink,
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

func (u *User) ExternalLink() string {
	return u.externalLink
}

func (u *User) Imege() string {
	return u.image
}

func (u *User) CreatedAt() vo.CreatedAt {
	return u.createdAt
}

func (u *User) UpdatedAt() vo.UpdatedAt {
	return u.updatedAt
}

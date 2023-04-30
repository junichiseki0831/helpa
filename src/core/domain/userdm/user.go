package domain

import (
	"errors"
	"helpa/src/core/domain/shared/vo"
	"time"
	"unicode/utf8"
)

type User struct {
	id           UserID
	name         string
	password     vo.Password
	email        vo.Email
	introduction string
	note         string
	image        vo.Image
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
	image vo.Image,
	createdAt vo.CreatedAt,
	updatedAt vo.UpdatedAt,
) (*User, error) {

	if name == "" {
		return nil, errors.New("empty name")
	} else if nameLen := utf8.RuneCountInString(name); nameLen > 50 {
		return nil, errors.New("please enter your name within 50 characters")
	}

	if introductionLen := utf8.RuneCountInString(introduction); introductionLen > 500 {
		return nil, errors.New("please enter your introduction within 500 characters")
	}

	if noteLen := utf8.RuneCountInString(note); noteLen > 500 {
		return nil, errors.New("please enter your note within 500 characters")
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

func (u *User) Image() vo.Image {
	return u.image
}

func (u *User) CreatedAt() vo.CreatedAt {
	return u.createdAt
}

func (u *User) UpdatedAt() vo.UpdatedAt {
	return u.updatedAt
}

func Reconstruct(id, name, password, email, introduction, note, image string, createdAt, updatedAt time.Time) (*User, error) {
	userID, err := NewUserIDByVal(id)
	if err != nil {
		return nil, err
	}
	pass, err := vo.NewPassword(password)
	if err != nil {
		return nil, err
	}
	mail, err := vo.NewEmail(email)
	if err != nil {
		return nil, err
	}
	img, err := vo.NewImage(image)
	if err != nil {
		return nil, err
	}
	ca, err := vo.NewCreatedAtByVal(createdAt)
	if err != nil {
		return nil, err
	}
	ua, err := vo.NewUpdatedAtByVal(updatedAt)
	if err != nil {
		return nil, err
	}

	return newUser(
		userID,
		name,
		pass,
		mail,
		introduction,
		note,
		img,
		ca,
		ua,
	)
}

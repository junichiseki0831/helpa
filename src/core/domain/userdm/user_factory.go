package domain

import (
	"encoding/base64"
	"helpa/src/core/domain/shared/vo"
	"time"
)

func GenForTest(id, name, password, email, introduction, note, image string, createdAt, updatedAt time.Time) (*User, error) {
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
	return newUser(userID, name, pass, mail, introduction, note, img, ca, ua)
}

func GenWhenCreate(name string, password vo.Password, email vo.Email, introduction string, note string, image []byte) (*User, error) {
	base64String := base64.StdEncoding.EncodeToString(image)
	img, err := vo.NewImage(base64String)
	if err != nil {
		return nil, err
	}
	return newUser(NewUserID(), name, password, email, introduction, note, img, vo.NewCreatedAt(), vo.NewUpdatedAt())
}

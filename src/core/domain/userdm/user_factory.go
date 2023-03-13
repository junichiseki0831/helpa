package domain

import (
	"fmt"
	"helpa/src/core/domain/shared/vo"
	"time"
)

func GenForTest(id, name, password, email, introduction, note, externalLink, imege string, createdAt, updatedAt time.Time) (*User, error) {
	userID, err := NewUserIDByVal(id)
	fmt.Println(userID)
	if err != nil {
		return nil, err
	}
	pass, err := vo.NewPassword(password)
	if err != nil {
		return nil, err
	}
	fmt.Println(pass)
	mail, err := vo.NewEmail(email)
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
	return NewUser(userID, name, pass, mail, introduction, note, externalLink, imege, ca, ua), nil
}

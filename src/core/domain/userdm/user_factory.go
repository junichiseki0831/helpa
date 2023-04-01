package domain

import (
	"helpa/src/core/domain/shared/vo"
	"time"
)

func GenForTest(id, name, password, email, introduction, note string, image []byte, createdAt, updatedAt time.Time) (*User, error) {
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
	ca, err := vo.NewCreatedAtByVal(createdAt)
	if err != nil {
		return nil, err
	}
	ua, err := vo.NewUpdatedAtByVal(updatedAt)
	if err != nil {
		return nil, err
	}
	return newUser(userID, name, pass, mail, introduction, note, image, ca, ua)
}

package domain

import (
	"github.com/google/uuid"
)

type UserID string

func GenerateUserID() UserID {
	ui := UserID(uuid.New().String())
	return ui
}

func (userID UserID) Equal(otherUserID UserID) bool {
	return userID == otherUserID
}

func (userID UserID) String() string {
	return string(userID)
}

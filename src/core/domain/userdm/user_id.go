package domain

import (
	"github.com/google/uuid"
)

type UserID string

func GenerateUserId() UserID {
	ui := UserID(uuid.New().String())
	return ui
}

func (userId UserID) Equal(otherUserId UserID) bool {
	return userId == otherUserId
}

func (userId UserID) String() string {
	return string(userId)
}

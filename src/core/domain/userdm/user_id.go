package domain

import (
	"errors"

	"github.com/google/uuid"
)

type UserID string

func NewUserID() UserID {
	return UserID(uuid.New().String())
}

func NewUserIDByVal(val string) (UserID, error) {
	if val == "" {
		return UserID(""), errors.New("user id must not be empty")
	}
	return UserID(val), nil
}

func (userID UserID) Equal(otherUserID UserID) bool {
	return userID == otherUserID
}

func (userID UserID) String() string {
	return string(userID)
}

package domain

import (
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

type UserID string

func NewUserID() UserID {
	ui := UserID(uuid.New().String())
	return ui
}

func NewUserIDByVal(val string) (UserID, error) {
	if val == "" {
		return UserID(""), xerrors.New("userdm id must not be empty")
	}
	return UserID(val), nil
}

func (userID UserID) Equal(otherUserID UserID) bool {
	return userID == otherUserID
}

func (userID UserID) String() string {
	return string(userID)
}

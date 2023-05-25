package domain

import (
	"helpa/src/support/smperr"

	"github.com/google/uuid"
)

type UserID string

func NewUserID() UserID {
	return UserID(uuid.New().String())
}

func NewUserIDByVal(val string) (UserID, error) {
	if val == "" {
		return UserID(""), smperr.BadRequest("user id must not be empty")
	}

	_, err := uuid.Parse(val)
	if err != nil {
		return UserID(""), smperr.BadRequest("invalid UUID length")
	}

	return UserID(val), nil
}

func (userID UserID) Equal(otherUserID UserID) bool {
	return userID == otherUserID
}

func (userID UserID) String() string {
	return string(userID)
}

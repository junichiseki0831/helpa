package domain

import (
    "github.com/google/uuid"
	"strings"
)

type UserID string

func GenerateUserId() string {
	uuid := uuid.New()
	id := ChangeStringType(uuid)
	return id
}

func ChangeStringType(uuid uuid.UUID) string {
	return strings.Replace(uuid.String(), "-", "", -1)
}

func (userId *UserID) Equal(otherUserId *UserID) bool {
	return userId == otherUserId
}

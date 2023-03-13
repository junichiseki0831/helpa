package domain_test

import (
	domain "helpa/src/core/domain/userdm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserIDToString(t *testing.T) {
	id := domain.NewUserID()
	result := id.String()
	assert.NotEmpty(t, result)
}

func TestEqual(t *testing.T) {
	id := domain.UserID("123-456-7890")
	result := id.Equal(id)
	assert.True(t, result)
}

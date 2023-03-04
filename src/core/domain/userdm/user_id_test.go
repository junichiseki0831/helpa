package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserIDToString(t *testing.T) {
	id := NewUserID()
	result := id.String()
	assert.NotEmpty(t, result)
}

func TestEqual(t *testing.T) {
	id := UserID("123-456-7890")
	result := id.Equal(id)
	assert.True(t, result)
}

package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUserIDToString(t *testing.T) {
	id := GenerateUserID()
	result := id.String()
	assert.NotEmpty(t, result)
}

func TestEqual(t *testing.T) {
	id := UserID("123-456-7890")
	result := id.Equal(id)
	assert.True(t, result)
}

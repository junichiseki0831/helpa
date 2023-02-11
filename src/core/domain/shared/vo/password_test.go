package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfirmPassword(t *testing.T) {
	password := Password("12345678")
	assert.NotEmpty(t, password)
}

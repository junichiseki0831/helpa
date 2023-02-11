package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfirmEmail(t *testing.T) {
	email := Email("test@test.co.jp")
	assert.NotEmpty(t, email)
}

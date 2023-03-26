package domain_test

import (
	domain "helpa/src/core/domain/userdm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserID(t *testing.T) {
	id := domain.NewUserID()
	result := id.String()
	assert.NotEmpty(t, result)
}

package domain_test

import (
	domain "helpa/src/core/domain/userdm"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, _ := domain.GenForTest("123", "testName", "12345671", "test1@test.com", "testIntroduction", "testNote", "testExternalLink", "image1", time.Now(), time.Now())

	assert.NotEmpty(t, user.ID())
	assert.NotEmpty(t, user.Name())
	assert.NotEmpty(t, user.Password())
	assert.NotEmpty(t, user.Email())
	assert.NotEmpty(t, user.Introduction())
	assert.NotEmpty(t, user.Note())
	assert.NotEmpty(t, user.ExternalLink())
	assert.NotEmpty(t, user.Imege())
	assert.NotEmpty(t, user.CreatedAt())
	assert.NotEmpty(t, user.UpdatedAt())
}

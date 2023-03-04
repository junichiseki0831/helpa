package domain_test

import (
	"helpa/src/core/domain/shared/vo"
	domain "helpa/src/core/domain/userdm"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestID(t *testing.T) {
	c, _ := vo.NewCreatedAt()
	u, _ := vo.NewUpdatedAt()
	i := domain.NewUserID()
	user := domain.NewUser(i, "test1", "12345671", "test1@test.com", "test1", "test1", "test1", "image1", c, u)

	assert.NotEmpty(t, user.ID())
}

func TestName(t *testing.T) {
	c, _ := vo.NewCreatedAt()
	u, _ := vo.NewUpdatedAt()
	i := domain.NewUserID()
	user := domain.NewUser(i, "test1", "12345671", "test1@test.com", "test1", "test1", "test1", "image1", c, u)

	assert.NotEmpty(t, user.Name())
}

func TestPassword(t *testing.T) {
	c, _ := vo.NewCreatedAt()
	u, _ := vo.NewUpdatedAt()
	i := domain.NewUserID()
	user := domain.NewUser(i, "test1", "12345671", "test1@test.com", "test1", "test1", "test1", "image1", c, u)

	assert.NotEmpty(t, user.Password())
}

func TestEmail(t *testing.T) {
	c, _ := vo.NewCreatedAt()
	u, _ := vo.NewUpdatedAt()
	i := domain.NewUserID()
	user := domain.NewUser(i, "test1", "12345671", "test1@test.com", "test1", "test1", "test1", "image1", c, u)

	assert.NotEmpty(t, user.Email())
}

func TestIntroduction(t *testing.T) {
	c, _ := vo.NewCreatedAt()
	u, _ := vo.NewUpdatedAt()
	i := domain.NewUserID()
	user := domain.NewUser(i, "test1", "12345671", "test1@test.com", "test1", "test1", "test1", "image1", c, u)

	assert.NotEmpty(t, user.Introduction())
}

func TestNote(t *testing.T) {
	c, _ := vo.NewCreatedAt()
	u, _ := vo.NewUpdatedAt()
	i := domain.NewUserID()
	user := domain.NewUser(i, "test1", "12345671", "test1@test.com", "test1", "test1", "test1", "image1", c, u)

	assert.NotEmpty(t, user.Note())
}

func TestExternalLink(t *testing.T) {
	c, _ := vo.NewCreatedAt()
	u, _ := vo.NewUpdatedAt()
	i := domain.NewUserID()
	user := domain.NewUser(i, "test1", "12345671", "test1@test.com", "test1", "test1", "test1", "image1", c, u)

	assert.NotEmpty(t, user.ExternalLink())
}

func TestImege(t *testing.T) {
	c, _ := vo.NewCreatedAt()
	u, _ := vo.NewUpdatedAt()
	i := domain.NewUserID()
	user := domain.NewUser(i, "test1", "12345671", "test1@test.com", "test1", "test1", "test1", "image1", c, u)

	assert.NotEmpty(t, user.Imege())
}

func TestCreatedAt(t *testing.T) {
	c, _ := vo.NewCreatedAt()
	u, _ := vo.NewUpdatedAt()
	i := domain.NewUserID()
	user := domain.NewUser(i, "test1", "12345671", "test1@test.com", "test1", "test1", "test1", "image1", c, u)

	assert.NotEmpty(t, user.CreatedAt())
}

func TestUpdatedAt(t *testing.T) {
	c, _ := vo.NewCreatedAt()
	u, _ := vo.NewUpdatedAt()
	i := domain.NewUserID()
	user := domain.NewUser(i, "test1", "12345671", "test1@test.com", "test1", "test1", "test1", "image1", c, u)

	assert.NotEmpty(t, user.UpdatedAt())
}

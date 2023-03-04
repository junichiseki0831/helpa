package vo_test

import (
	"helpa/src/core/domain/shared/vo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmail(t *testing.T) {
	cases := map[int]string{
		1: "test1@test.co.jp",
		2: "test2@test.co.jp",
		3: "test3@test.co.jp",
		4: "test4@test.co.jp",
		5: "test5@test.co.jp",
	}

	got, err := vo.NewEmail(cases[1])
	assert.Empty(t, err)
	assert.Equal(t, "test1@test.co.jp", got.String())
}

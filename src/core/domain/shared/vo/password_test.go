package vo_test

import (
	"helpa/src/core/domain/shared/vo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPassword(t *testing.T) {
	cases := map[int]string{
		1: "12345671",
		2: "12345672",
		3: "12345673",
		4: "12345674",
		5: "12345675",
	}

	got, err := vo.NewPassword(cases[1])
	assert.Empty(t, err)
	assert.Equal(t, "12345671", got.String())
}

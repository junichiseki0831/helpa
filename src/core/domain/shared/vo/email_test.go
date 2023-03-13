package vo_test

import (
	"helpa/src/core/domain/shared/vo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmail(t *testing.T) {
	cases1 := map[string]string{
		"正常形1": "test1@test.co.jp",
		"正常形2": "test2@test.co.jp",
		"正常形3": "test3@test.co.jp",
	}

	cases2 := map[string]string{
		"異常形1": "test4test.co.jp",
		"異常形2": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@aaa.com",
	}

	for name, tt := range cases1 {
		t.Run(name, func(t *testing.T) {
			got, err := vo.NewEmail(tt)
			assert.Empty(t, err)
			assert.Equal(t, tt, got.String())
		})
	}

	for name, tt := range cases2 {
		t.Run(name, func(t *testing.T) {
			_, err := vo.NewEmail(tt)
			assert.Error(t, err)
		})
	}
}

package vo_test

import (
	"helpa/src/core/domain/shared/vo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPassword(t *testing.T) {
	cases1 := map[string]string{
		"正常形1": "12345671",
		"正常形2": "12345672",
		"正常形3": "12345673",
	}

	cases2 := map[string]string{
		"異常形1": "123456789",
		"異常形2": "1234567@",
	}

	for name, tt := range cases1 {
		t.Run(name, func(t *testing.T) {
			got, err := vo.NewPassword(tt)
			assert.Empty(t, err)
			assert.Equal(t, tt, got.String())
		})
	}

	for name, tt := range cases2 {
		t.Run(name, func(t *testing.T) {
			_, err := vo.NewPassword(tt)
			assert.Error(t, err)
		})
	}
}

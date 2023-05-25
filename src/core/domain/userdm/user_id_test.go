package domain_test

import (
	domain "helpa/src/core/domain/userdm"
	"helpa/src/support/smperr"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserID(t *testing.T) {
	id := domain.NewUserID()
	result := id.String()
	assert.NotEmpty(t, result)
}

func TestNewUserIDByVal(t *testing.T) {
	for _, tt := range []struct {
		name  string
		input string
		isErr bool
	}{
		{
			name:  "正常系",
			input: "7b8a6b7a-3831-4f57-94c9-29f0df3a68f9",
		},
		{
			name:  "異常系: 引数が空文字",
			input: "",
			isErr: true,
		},
		{
			name:  "異常系: 引数の文字列が無効なUUIDの長さ",
			input: "aabbddd",
			isErr: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := domain.NewUserIDByVal(tt.input)
			if tt.isErr {
				assert.IsType(t, &smperr.BadRequestErr{}, err)
				assert.Empty(t, got)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.input, got.String())
		})
	}
}

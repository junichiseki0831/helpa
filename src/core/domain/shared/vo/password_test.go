package vo_test

import (
	"helpa/src/core/domain/shared/vo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPassword(t *testing.T) {
	for _, tt := range []struct {
		name  string
		input string
		isErr bool
	}{
		{
			name:  "正常系",
			input: "12345671",
		},
		{
			name:  "異常系: 引数が8文字以上",
			input: "123456789",
			isErr: true,
		},
		{
			name:  "異常系: 引数に記号が含まれる",
			input: "1234567@",
			isErr: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vo.NewPassword(tt.input)
			if tt.isErr {
				assert.NotNil(t, err)
				assert.Empty(t, got)
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, tt.input, got.String())
		})
	}
}

package vo_test

import (
	"helpa/src/core/domain/shared/vo"
	"helpa/src/support/smperr"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmail(t *testing.T) {
	for _, tt := range []struct {
		name  string
		input string
		isErr bool
	}{
		{
			name:  "正常系",
			input: "test@test.co.jp",
		},
		{
			name:  "異常系: 引数が256文字以上",
			input: strings.Repeat("a", 256) + "@a.com",
			isErr: true,
		},
		{
			name:  "異常系: 引数が@を含まない",
			input: "test4test.co.jp",
			isErr: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vo.NewEmail(tt.input)
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

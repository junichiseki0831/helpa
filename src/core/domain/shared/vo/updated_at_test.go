package vo_test

import (
	"helpa/src/core/domain/shared/vo"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// NewUpdatedAt で値を作成し、先に作成した値よりも後に作成されている値かを確認する
func TestNewUpdatedAt(t *testing.T) {
	for _, tt := range []struct {
		name  string
		input time.Time
		isErr bool
	}{
		{
			name:  "正常系",
			input: time.Now(),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := vo.NewUpdatedAt()
			assert.True(t, got.Value().After(tt.input))
		})
	}
}

func TestNewUpdatedAtByVal(t *testing.T) {
	for _, tt := range []struct {
		name  string
		input time.Time
		isErr bool
	}{
		{
			name:  "正常系",
			input: time.Date(2023, time.December, 10, 23, 1, 10, 0, time.Local),
		},
		{
			name:  "異常系: 引数がtime.IsZero",
			input: time.Time{},
			isErr: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got, err := vo.NewUpdatedAtByVal(tt.input)
			if tt.isErr {
				assert.NotNil(t, err)
				assert.Empty(t, got)
				return
			}

			assert.Nil(t, err)
			assert.True(t, got.Value().Equal(tt.input))
		})
	}
}

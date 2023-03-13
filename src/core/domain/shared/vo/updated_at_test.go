package vo_test

import (
	"helpa/src/core/domain/shared/vo"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// NewUpdatedAt で値を作成し、先に作成した値よりも後に作成されている値かを確認する
func TestNewUpdatedAt(t *testing.T) {
	cases1 := map[string]time.Time{
		"正常形1": time.Now(),
		"正常形2": time.Now(),
		"正常形3": time.Now(),
	}
	cases2 := map[string]time.Time{
		"異常形1": time.Date(2023, time.December, 10, 23, 1, 10, 0, time.Local),
		"異常形2": time.Date(2023, time.December, 10, 23, 1, 10, 0, time.Local),
		"異常形3": time.Date(2023, time.December, 10, 23, 1, 10, 0, time.Local),
	}

	for name, tt := range cases1 {
		t.Run(name, func(t *testing.T) {
			got, err := vo.NewUpdatedAt()
			assert.Empty(t, err)
			assert.True(t, got.Value().After(tt))
		})
	}

	for name, tt := range cases2 {
		t.Run(name, func(t *testing.T) {
			got, err := vo.NewUpdatedAt()
			assert.Empty(t, err)
			assert.False(t, got.Value().After(tt))
		})
	}
}

// 同じ値を NewUpdatedAtByVal で作成し、同じ内容を返し比較する
func TestUpdatedAtEqual(t *testing.T) {
	cases := map[int]time.Time{
		1: time.Date(2023, time.December, 10, 23, 1, 10, 0, time.Local),
		2: time.Date(2023, time.December, 10, 23, 1, 10, 0, time.Local),
	}
	got1, err1 := vo.NewUpdatedAtByVal(cases[1])
	got2, err2 := vo.NewUpdatedAtByVal(cases[2])
	result := got1.Equal(got2)
	assert.Empty(t, err1, err2)
	assert.True(t, result)
}
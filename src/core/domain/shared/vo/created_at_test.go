package vo_test

import (
	"helpa/src/core/domain/shared/vo"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// NewCreatedAt で値を作成し、先に作成した値よりも後に作成されている値かを確認する
func TestNewCreatedAt(t *testing.T) {
	cases := map[int]time.Time{
		1: time.Now(),
		2: time.Date(2023, time.December, 10, 23, 1, 10, 0, time.Local),
		3: time.Date(2023, time.December, 10, 23, 1, 10, 0, time.Local),
		4: time.Now(),
		5: time.Now(),
	}
	got, err := vo.NewCreatedAt()
	assert.Empty(t, err)
	assert.True(t, got.Value().After(cases[1]))
}

// 同じ値を NewCreatedAtByVal で作成し、同じ内容を返し比較する
func TestEqual(t *testing.T) {
	cases := map[int]time.Time{
		1: time.Now(),
		2: time.Date(2023, time.December, 10, 23, 1, 10, 0, time.Local),
		3: time.Date(2023, time.December, 10, 23, 1, 10, 0, time.Local),
		4: time.Now(),
		5: time.Now(),
	}
	got1, err1 := vo.NewCreatedAtByVal(cases[2])
	got2, err2 := vo.NewCreatedAtByVal(cases[3])
	result := got1.Equal(got2)
	assert.Empty(t, err1, err2)
	assert.True(t, result)
}

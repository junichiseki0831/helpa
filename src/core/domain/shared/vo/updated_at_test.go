package vo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUpdatedAt(t *testing.T) {
	now := time.Now()
	assert.NotEmpty(t, UpdatedAt(now.Format("2006-01-02 15:04:05")))
}

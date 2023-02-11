package vo

import (
	"time"
)

type CreatedAt string

const DateTimeFormat = "2006-01-02 15:04:05"

func GenerateCreatedAt() CreatedAt {
	t := time.Now()
	return CreatedAt(t.Format(DateTimeFormat))
}

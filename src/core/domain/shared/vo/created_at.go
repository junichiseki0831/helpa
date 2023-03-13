package vo

import (
	"time"
)

type CreatedAt time.Time

func NewCreatedAt() (CreatedAt, error) {
	t := CreatedAt(time.Now())
	return t, nil
}

func (c CreatedAt) Value() time.Time {
	return time.Time(c)
}

func NewCreatedAtByVal(at time.Time) (CreatedAt, error) {
	t := CreatedAt(at)
	return t, nil
}

func (c CreatedAt) String(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func (c CreatedAt) Equal(c2 CreatedAt) bool {
	return c == c2
}

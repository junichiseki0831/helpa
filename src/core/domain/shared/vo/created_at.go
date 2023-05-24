package vo

import (
	"helpa/src/support/smperr"
	"time"
)

type CreatedAt time.Time

func NewCreatedAt() CreatedAt {
	return CreatedAt(time.Now())
}

func (c CreatedAt) Value() time.Time {
	return time.Time(c)
}

func NewCreatedAtByVal(at time.Time) (CreatedAt, error) {
	t := CreatedAt(at)
	if at.IsZero() {
		return CreatedAt(time.Time{}), smperr.BadRequest("invalid time provided")
	}
	return t, nil
}

func (c CreatedAt) String() string {
	return c.Value().Format("2006-01-02 15:04:05")
}

func (c CreatedAt) Equal(c2 CreatedAt) bool {
	return c == c2
}

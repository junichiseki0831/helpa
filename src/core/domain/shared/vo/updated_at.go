package vo

import (
	"errors"
	"time"
)

type UpdatedAt time.Time

func NewUpdatedAt() (UpdatedAt, error) {
	t := UpdatedAt(time.Now())
	return t, nil
}

func (u UpdatedAt) Value() time.Time {
	return time.Time(u)
}

func NewUpdatedAtByVal(ut time.Time) (UpdatedAt, error) {
	t := UpdatedAt(ut)
	if ut.IsZero() {
		return UpdatedAt(time.Time{}), errors.New("error")
	}
	return t, nil
}

func (u UpdatedAt) String(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func (u UpdatedAt) Equal(u2 UpdatedAt) bool {
	return u == u2
}

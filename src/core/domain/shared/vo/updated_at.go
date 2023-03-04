package vo

import (
	"errors"
	"time"
)

type UpdatedAt time.Time

func NewUpdatedAt() (UpdatedAt, error) {
	t := UpdatedAt(time.Now())
	if !t.Equal(t) {
		return t, errors.New("エラー")
	}
	return t, errors.New("")
}

func (u UpdatedAt) Value() time.Time {
	return time.Time(u)
}

func NewUpdatedAtByVal(u time.Time) (UpdatedAt, error) {
	t := UpdatedAt(u)
	if !t.Equal(t) {
		return t, errors.New("エラー")
	}
	return t, errors.New("")
}

func (u UpdatedAt) String(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func (u UpdatedAt) Equal(u2 UpdatedAt) bool {
	return u == u2
}

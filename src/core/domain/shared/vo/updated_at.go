package vo

import (
	"helpa/src/support/smperr"
	"time"
)

type UpdatedAt time.Time

func NewUpdatedAt() UpdatedAt {
	return UpdatedAt(time.Now())
}

func (u UpdatedAt) Value() time.Time {
	return time.Time(u)
}

func NewUpdatedAtByVal(ut time.Time) (UpdatedAt, error) {
	t := UpdatedAt(ut)
	if ut.IsZero() {
		return UpdatedAt(time.Time{}), smperr.BadRequest("invalid time provided")
	}
	return t, nil
}

func (u UpdatedAt) String() string {
	return u.Value().Format("2006-01-02 15:04:05")
}

func (u UpdatedAt) Equal(u2 UpdatedAt) bool {
	return u == u2
}

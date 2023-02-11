package vo

import "time"

type UpdatedAt string

const UpdatedAtDateTimeFormat = "2006-01-02 15:04:05"

func GenerateUpdatedAt() UpdatedAt {
	t := time.Now()
	return UpdatedAt(t.Format(UpdatedAtDateTimeFormat))
}

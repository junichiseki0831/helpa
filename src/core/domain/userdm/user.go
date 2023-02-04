package domain

import "time"
type User struct {
  id UserID
  name string
  password string
  email string
  introduction string
  note string
  external_link string
  imege string
  CreatedAt time.Time
  UpdatedAt time.Time
}

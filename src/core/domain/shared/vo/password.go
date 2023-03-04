package vo

import "errors"

type Password string

func NewPassword(password string) (Password, error) {
	len := len(password)
	if len < 8 || len > 8 {
		return Password(""), errors.New("please use 8 characters")
	}
	return Password(password), errors.New("")
}

func (p Password) String() string {
	return string(p)
}

func (p Password) Equal(p2 Password) bool {
	return p == p2
}

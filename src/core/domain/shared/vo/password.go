package vo

import (
	"errors"
	"regexp"
)

type Password string

func NewPassword(password string) (Password, error) {
	len := len(password)
	if len < 8 || len > 8 { // 文字数判定
		return Password(""), errors.New("please use 8 characters")
	}
	if !(regexp.MustCompile("^[0-9a-zA-Z]+$").MatchString(password)) { // 英数字以外を使っているか判定
		return Password(""), errors.New("please do not use non-alphanumeric characters")
	}
	return Password(password), nil
}

func (p Password) String() string {
	return string(p)
}

func (p Password) Equal(p2 Password) bool {
	return p == p2
}

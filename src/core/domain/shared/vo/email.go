package vo

import (
	"errors"
	"regexp"
)

type Email string

var emailCheckFormat = regexp.MustCompile(`^(?i:[^ @"<>]+|".*")@(?i:[a-z1-9.])+.(?i:[a-z])+$`)

func NewEmail(email string) (Email, error) {
	len := len(email)
	if len >= 256 {
		return Email(""), errors.New("please enter within 256 characters")
	}
	if !emailCheckFormat.MatchString(email) {
		return Email(""), errors.New("please fill in according to the email format")
	}
	return Email(email), nil
}

func (e Email) String() string {
	return string(e)
}

func (e Email) Equal(e2 Email) bool {
	return e == e2
}

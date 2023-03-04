package vo

import "errors"

type Email string

func NewEmail(email string) (Email, error) {
	len := len(email)
	if len <= 120 {
		return Email(email), errors.New("")
	}
	return Email(""), errors.New("エラー")
}

func (e Email) String() string {
	return string(e)
}

func (e Email) Equal(e2 Email) bool {
	return e == e2
}

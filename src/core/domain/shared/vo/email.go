package vo

import "errors"

type Email string

func ConfirmEmail(email string) (Email error) {
	len := len(email)
	if len <= 120 {
		return errors.New("please enter within 120 characters")
	} else {
		return Email
	}
}

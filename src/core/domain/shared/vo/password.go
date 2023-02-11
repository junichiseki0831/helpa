package vo

import "errors"

type Password string

func ConfirmPassword(password string) (Password error) {
	len := len(password)
	if len < 8 || len > 8 {
		return errors.New("please use 8 characters")
	} else {
		return Password
	}
}

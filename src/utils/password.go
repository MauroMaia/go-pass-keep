package utils

import "errors"

func ValidatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("Invalid Password len. Must be >= 6")
	}
	return nil
}

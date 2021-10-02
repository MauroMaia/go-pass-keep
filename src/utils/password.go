package utils

import "errors"

func ValidatePassword(password string) error {
	if len(password) > 8 {
		return errors.New("Invalid Password len. Must be >= 8")
	}
	return nil
}

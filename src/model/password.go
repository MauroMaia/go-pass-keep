package model

import (
	"go-pass-keeper/src/utils"
)

type Password struct {
	password string
}

func NewPassword(password string) (*Password, error) {
	p := new(Password)
	return p.SetPassword(password)
}

func (p Password) GetPassword() (string, error) {
	return p.password, nil
}

func (p Password) SetPassword(newPassword string) (*Password, error) {
	if err := utils.ValidatePassword(newPassword); err != nil {
		return &p, err
	}

	// Just for user information
	utils.IsPasswordWeak(newPassword)

	p.password = newPassword
	return &p, nil
}

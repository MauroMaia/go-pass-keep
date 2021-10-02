package model

import (
	"errors"
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
	if len(newPassword) > 8 {
		p.password = newPassword
		return &p, nil
	}
	return &p, errors.New("Password len is invalid. Must be >= 8")
}

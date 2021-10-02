package model

import (
	"encoding/json"
	"time"
)

type Entry struct {
	// id string
	title     string
	username  string
	password  *Password
	createdAt string
	updatedAt string
}

func NewEntry(title string, user string, password string) (*Entry, error) {
	p, err := NewPassword(password)
	if err != nil {
		return nil, err
	}

	e := &Entry{
		title:     title,
		username:  user,
		password:  p,
		createdAt: time.Now().Format(time.RFC3339),
		updatedAt: time.Now().Format(time.RFC3339),
	}

	return e, nil
}

func (m Entry) MarshalJSON() ([]byte, error) {
	password, err := m.password.GetPassword()
	if err != nil {
		return nil, err
	}

	return json.Marshal(struct {
		Title     string `json:"title"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		m.title,
		m.username,
		password,
		m.createdAt,
		m.updatedAt,
	})
}

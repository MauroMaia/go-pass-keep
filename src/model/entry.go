package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Entry struct {
	uuid      string
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
		uuid:      uuid.Must(uuid.NewRandom()).String(),
		title:     title,
		username:  user,
		password:  p,
		createdAt: time.Now().Format(time.RFC3339),
		updatedAt: time.Now().Format(time.RFC3339),
	}

	return e, nil
}

func NewEntryFull(uuid string, title string, user string, password string, createdAt string, updatedAt string) (*Entry, error) {
	p, err := NewPassword(password)
	if err != nil {
		return nil, err
	}

	e := &Entry{
		uuid:      uuid,
		title:     title,
		username:  user,
		password:  p,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}

	return e, nil
}

func NewEntryFromMap(mapStruct map[string]interface{}) (*Entry, error) {
	return NewEntryFull(
		mapStruct["uuid"].(string),
		mapStruct["title"].(string),
		mapStruct["username"].(string),
		mapStruct["password"].(string),
		mapStruct["created_at"].(string),
		mapStruct["updated_at"].(string),
	)
}

func (m Entry) MarshalJSON() ([]byte, error) {
	password, err := m.password.GetPassword()
	if err != nil {
		return nil, err
	}

	return json.Marshal(struct {
		Uuid      string `json:"uuid"`
		Title     string `json:"title"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}{
		m.uuid,
		m.title,
		m.username,
		password,
		m.createdAt,
		m.updatedAt,
	})
}

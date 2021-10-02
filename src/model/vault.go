package model

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type Vault struct {
	uuid      string
	title     string
	entries   []*Entry
	createdAt string
	updatedAt string
}

func NewVault(title string) (*Vault, error) {
	if title == "" {
		return nil, errors.New("Vault title must not be empty")
	}

	v := &Vault{
		uuid:      uuid.Must(uuid.NewRandom()).String(),
		title:     title,
		createdAt: time.Now().Format(time.RFC3339),
		updatedAt: time.Now().Format(time.RFC3339),
	}

	return v, nil
}

func (m Vault) MarshalJSON() ([]byte, error) {

	return json.Marshal(struct {
		Uuid      string   `json:"uuid"`
		Title     string   `json:"title"`
		Entries   []*Entry `json:"entries"`
		CreatedAt string   `json:"created_at"`
		UpdatedAt string   `json:"updated_at"`
	}{
		m.uuid,
		m.title,
		m.entries,
		m.createdAt,
		m.updatedAt,
	})
}

func (vault *Vault) UnmarshalJSON(data []byte) (err error) {
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	keys := reflect.ValueOf(result).MapKeys()
	for _, item := range keys {
		keyname := item.Interface().(string)

		if keyname == "uuid" {
			vault.uuid = result[keyname].(string)
		} else if keyname == "title" {
			vault.title = result[keyname].(string)
		} else if keyname == "created_at" {
			vault.createdAt = result[keyname].(string)
		} else if keyname == "updated_at" {
			vault.updatedAt = result[keyname].(string)
		} else if keyname == "entries" && result[keyname] != nil {

			for _, entry := range result[keyname].([]interface{}) {

				realEntry, err := NewEntryFromMap(entry.(map[string]interface{}))
				if err != nil {
					return err
				}
				vault.entries = append(vault.entries, realEntry)
			}
		}
	}
	return nil
}

func (vault Vault) ContainsEntry(user string, title string) bool {

	for _, entry := range vault.entries {
		if entry.title == title && entry.username == user {
			return true
		}
	}
	return false
}

func (vault Vault) GetAllEntries() []*Entry {
	return vault.entries
}

func (vault Vault) PutEntry(entry *Entry) *Vault {
	vault.entries = append(vault.entries, entry)
	vault.updatedAt = time.Now().Format(time.RFC3339)
	return &vault
}

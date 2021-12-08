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
	groups    []*Group
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

func (vault Vault) MarshalJSON() ([]byte, error) {

	return json.Marshal(struct {
		Uuid      string   `json:"uuid"`
		Title     string   `json:"title"`
		Entries   []*Entry `json:"entries"`
		Groups    []*Group `json:"groups"`
		CreatedAt string   `json:"created_at"`
		UpdatedAt string   `json:"updated_at"`
	}{
		vault.uuid,
		vault.title,
		vault.entries,
		vault.groups,
		vault.createdAt,
		vault.updatedAt,
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
		} else if keyname == "groups" && result[keyname] != nil {

			for _, entry := range result[keyname].([]interface{}) {

				realGroup, err := NewGroupFromMap(entry.(map[string]interface{}))
				if err != nil {
					return err
				}
				vault.groups = append(vault.groups, realGroup)
			}
		}
	}
	return nil
}

// ContainsEntry This method check if one entry exist give its username,
// tile/description, and groupPath/groupTree names. The search is performed in
// a top down way.
func (vault Vault) ContainsEntry(user string, title string, groupPath []string) bool {

	for _, entry := range vault.entries {
		if entry.title == title && entry.username == user {
			return true
		}
	}
	if groupPath != nil {
		var thisGroupName string

		if len(groupPath) == 1 {
			thisGroupName = groupPath[0]
			groupPath = nil
		} else {
			thisGroupName = groupPath[0]
			groupPath = groupPath[1:]
		}

		for _, group := range vault.groups {
			if thisGroupName == group.GetTitle() && group.ContainsEntry(user, title, groupPath) {
				return true
			}
		}
	}
	return false
}

func (vault Vault) GetAllEntries() []*Entry {
	return vault.entries
}

// PutEntryInVault This method append an entry to the vault creating any
// group/SubGroup if needed.
// WARNING: There is no check if the entry already exist.
func (vault Vault) PutEntryInVault(entry *Entry, groupPath []string) (Vault, error) {
	if groupPath == nil {
		vault.entries = append(vault.entries, entry)
		vault.updatedAt = time.Now().Format(time.RFC3339)
		return vault, nil
	}

	var thisGroupName string

	if len(groupPath) == 1 {
		thisGroupName = groupPath[0]
		groupPath = nil
	} else {
		thisGroupName = groupPath[0]
		groupPath = groupPath[1:]
	}

	for index, group := range vault.groups {
		if group.GetTitle() == thisGroupName {
			group, err := group.PutEntryInGroup(entry, groupPath)
			if err != nil {
				return vault, err
			}
			vault.groups[index] = group
			vault.updatedAt = time.Now().Format(time.RFC3339)
			return vault, err
		}
	}

	/*
	 * if the logic came to here the group does not exit and need to be created
	 */
	group, err := NewGroup(thisGroupName)
	if err != nil {
		return vault, err
	}

	group, err = group.PutEntryInGroup(entry, groupPath)
	if err != nil {
		return vault, err
	}

	vault.groups = append(vault.groups, group)
	vault.updatedAt = time.Now().Format(time.RFC3339)
	return vault, nil
}

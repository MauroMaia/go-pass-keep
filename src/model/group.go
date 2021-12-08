package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Group struct {
	uuid      string
	title     string
	entries   []*Entry
	groups    []*Group
	createdAt string
	updatedAt string
}

func (group Group) GetTitle() string {
	return group.title
}

func NewGroup(title string) (*Group, error) {

	e := &Group{
		uuid:      uuid.Must(uuid.NewRandom()).String(),
		title:     title,
		createdAt: time.Now().Format(time.RFC3339),
		updatedAt: time.Now().Format(time.RFC3339),
	}

	return e, nil
}

func NewGroupFull(uuid string, title string, entries []*Entry, groups []*Group, createdAt string, updatedAt string) (*Group, error) {

	e := &Group{
		uuid:      uuid,
		title:     title,
		entries:   entries,
		groups:    groups,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}

	return e, nil
}

// NewGroupFromMap this function was created to allow custom UnmarshalJSON
// functions to restore objects of type Groups
func NewGroupFromMap(mapStruct map[string]interface{}) (*Group, error) {
	var entries []*Entry
	if mapStruct["entries"] != nil {
		for _, entry := range mapStruct["entries"].([]interface{}) {

			realEntry, err := NewEntryFromMap(entry.(map[string]interface{}))
			if err != nil {
				return nil, err
			}
			entries = append(entries, realEntry)
		}
	}

	var groups []*Group
	if mapStruct["groups"] != nil {
		for _, entry := range mapStruct["groups"].([]interface{}) {

			realGroup, err := NewGroupFromMap(entry.(map[string]interface{}))
			if err != nil {
				return nil, err
			}
			groups = append(groups, realGroup)
		}
	}

	return NewGroupFull(
		mapStruct["uuid"].(string),
		mapStruct["title"].(string),
		entries,
		groups,
		mapStruct["created_at"].(string),
		mapStruct["updated_at"].(string),
	)
}

func (group Group) ContainsEntry(user string, title string, groupPath []string) bool {

	for _, entry := range group.entries {
		if entry.title == title && entry.username == user {
			return true
		}
	}
	return false
}

// PutEntryInGroup This method append an entry to this group, creating any
// group/SubGroup if needed.
// WARNING: There is no check if the entry already exist.
func (group Group) PutEntryInGroup(entry *Entry, groupPath []string) (*Group, error) {
	if groupPath == nil {
		group.entries = append(group.entries, entry)
		group.updatedAt = time.Now().Format(time.RFC3339)
		return &group, nil
	}

	var thisGroupName string

	if len(groupPath) == 1 {
		thisGroupName = groupPath[0]
		groupPath = nil
	} else {
		thisGroupName = groupPath[0]
		groupPath = groupPath[1:]
	}

	for index, subGroup := range group.groups {
		if subGroup.GetTitle() == thisGroupName {
			subGroup, err := subGroup.PutEntryInGroup(entry, groupPath)
			if err != nil {
				return nil, err
			}
			group.groups[index] = subGroup
			group.updatedAt = time.Now().Format(time.RFC3339)
			return &group, nil
		}
	}

	// if it get here the group does not exit and need to be created
	subGroup, err := NewGroup(thisGroupName)
	if err != nil {
		return nil, err
	}

	subGroup, err = subGroup.PutEntryInGroup(entry, groupPath)
	if err != nil {
		return nil, err
	}

	group.groups = append(group.groups, subGroup)
	group.updatedAt = time.Now().Format(time.RFC3339)

	return &group, nil
}

func (group Group) MarshalJSON() ([]byte, error) {

	return json.Marshal(struct {
		Uuid      string   `json:"uuid"`
		Title     string   `json:"title"`
		Entries   []*Entry `json:"entries"`
		Groups    []*Group `json:"groups"`
		CreatedAt string   `json:"created_at"`
		UpdatedAt string   `json:"updated_at"`
	}{
		group.uuid,
		group.title,
		group.entries,
		group.groups,
		group.createdAt,
		group.updatedAt,
	})
}

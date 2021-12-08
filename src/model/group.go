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
	createdAt string
	updatedAt string
}

func (g Group) GetTitle() string {
	return g.title
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

func NewGroupFull(uuid string, title string, entries []*Entry, createdAt string, updatedAt string) (*Group, error) {

	e := &Group{
		uuid:      uuid,
		title:     title,
		entries:   entries,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}

	return e, nil
}

func NewGroupFromMap(mapStruct map[string]interface{}) (*Group, error) {
	return NewGroupFull(
		mapStruct["uuid"].(string),
		mapStruct["title"].(string),
		mapStruct["entries"].([]*Entry),
		mapStruct["created_at"].(string),
		mapStruct["updated_at"].(string),
	)
}

func (g Group) MarshalJSON() ([]byte, error) {

	return json.Marshal(struct {
		Uuid      string   `json:"uuid"`
		Title     string   `json:"title"`
		Entries   []*Entry `json:"entries"`
		CreatedAt string   `json:"created_at"`
		UpdatedAt string   `json:"updated_at"`
	}{
		g.uuid,
		g.title,
		g.entries,
		g.createdAt,
		g.updatedAt,
	})
}

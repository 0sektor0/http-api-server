package models

import (
	"encoding/json"
)

type Forum struct {
	Id      int64  `json:"-"`
	Threads int32  `json:"threads,omitempty"`
	Posts   int64  `json:"posts,omitempty"`
	Title   string `json:"title"`
	Slug    string `json:"slug"`
	User    string `json:"user"`
}

func UnmarshalForum(b []byte) (*Forum, error) {
	unmarshaledStruct := new(Forum)

	if err := json.Unmarshal(b, unmarshaledStruct); err != nil {
		return nil, err
	}

	return unmarshaledStruct, nil
}

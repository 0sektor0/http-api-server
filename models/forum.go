package models

import (
	"encoding/json"
)

type Forum struct {
	Threads int32  `json:"threads"`
	Posts   int64  `json:"posts"`
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

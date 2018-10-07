package models

import (
	"encoding/json"
)

type User struct {
	Id       int64  `json:"-"`
	About    string `json:"about,omitempty"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Nickname string `json:"nickname,omitempty"`
}

func UnmarshalUser(b []byte) (*User, error) {
	unmarshaledStruct := new(User)

	if err := json.Unmarshal(b, unmarshaledStruct); err != nil {
		return nil, err
	}

	return unmarshaledStruct, nil
}

package models

import "encoding/json"

type User struct {
	About    string `json:"about"`
	Email    string `json:"email"`
	FullName string `json:"fullname"`
	NickName string `json:"nickname"`
}

func UnmarshalUser(b []byte) (*User, error) {
	unmarshaledStruct := new(User)

	if err := json.Unmarshal(b, unmarshaledStruct); err != nil {
		return nil, err
	}

	return unmarshaledStruct, nil
}

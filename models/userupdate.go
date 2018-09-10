package models

import "encoding/json"

type UserUpdate struct {
	About    string `json:"about"`
	Email    string `json:"email"`
	FullName string `json:"fullname"`
}

func UnmarshalUserUpdate(b []byte) (*UserUpdate, error) {
	unmarshaledStruct := new(UserUpdate)

	if err := json.Unmarshal(b, unmarshaledStruct); err != nil {
		return nil, err
	}

	return unmarshaledStruct, nil
}

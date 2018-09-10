package models

import "encoding/json"

type Status struct {
	User   int32 `json:"user"`
	Post   int64 `json:"post"`
	Forum  int32 `json:"forum"`
	Thread int32 `json:"thread"`
}

func UnmarshalStatus(b []byte) (*Status, error) {
	unmarshaledStruct := new(Status)

	if err := json.Unmarshal(b, unmarshaledStruct); err != nil {
		return nil, err
	}

	return unmarshaledStruct, nil
}

package models

import "encoding/json"

type ThreadUpdate struct {
	Message string `json:"message"`
	Title   string `json:"title"`
}

func UnmarshalThreadUpdate(b []byte) (*ThreadUpdate, error) {
	unmarshaledStruct := new(ThreadUpdate)

	if err := json.Unmarshal(b, unmarshaledStruct); err != nil {
		return nil, err
	}

	return unmarshaledStruct, nil
}

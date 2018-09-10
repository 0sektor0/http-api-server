package models

import (
	"encoding/json"
)

type Error struct {
	Message string `json:"message"`
}

func UnmarshalError(b []byte) (*Error, error) {
	unmarshaledStruct := new(Error)

	if err := json.Unmarshal(b, unmarshaledStruct); err != nil {
		return nil, err
	}

	return unmarshaledStruct, nil
}

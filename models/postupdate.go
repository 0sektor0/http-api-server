package models

import "encoding/json"

type PostUpdate struct {
	Message string `json:"message"`
}

func UnmarshalPostUpdate(b []byte) (*PostUpdate, error) {
	unmarshaledStruct := new(PostUpdate)

	if err := json.Unmarshal(b, unmarshaledStruct); err != nil {
		return nil, err
	}

	return unmarshaledStruct, nil
}

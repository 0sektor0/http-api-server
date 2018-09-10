package models

import "encoding/json"

type Vote struct {
	NickName string `json:"nickname"`
	Voice    int32  `json:"voice"`
}

func UnmarshalVote(b []byte) (*Vote, error) {
	unmarshaledStruct := new(Vote)

	if err := json.Unmarshal(b, unmarshaledStruct); err != nil {
		return nil, err
	}

	return unmarshaledStruct, nil
}

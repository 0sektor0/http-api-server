package models

import "encoding/json"

type Thread struct {
	Id      int32  `json:"id"`
	Slug    string `json:"slug",omitempty`
	Title   string `json:"title"`
	Votes   int32  `json:"votes"`
	Forum   string `json:"forum"`
	Author  string `json:"author"`
	Created string `json:"created",omitempty`
	Message string `json:"message"`
}

func UnmarshalThread(b []byte) (*Thread, error) {
	unmarshaledStruct := new(Thread)

	if err := json.Unmarshal(b, unmarshaledStruct); err != nil {
		return nil, err
	}

	return unmarshaledStruct, nil
}

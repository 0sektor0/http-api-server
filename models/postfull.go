package models

import "encoding/json"

type PostFull struct {
	Post   *Post   `json:"post"`
	Forum  *Forum  `json:"forum"`
	Author *User   `json:"author"`
	Thread *Thread `json:"thread"`
}

func UnmarshalPostFull(b []byte) (*PostFull, error) {
	unmarshaledStruct := new(PostFull)

	if err := json.Unmarshal(b, unmarshaledStruct); err != nil {
		return nil, err
	}

	return unmarshaledStruct, nil
}

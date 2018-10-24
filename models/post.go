package models

import "encoding/json"

type Post struct {
	Id       int    `json:"id"`
	Forum    string `json:"forum"`
	Parent   *int   `json:"parent"`
	Thread   int    `json:"thread"`
	Author   string `json:"author"`
	Message  string `json:"message"`
	Created  string `json:"created"`
	Isedited bool   `json:"isEdited"`
}

func UnmarshalPost(b []byte) (*Post, error) {
	unmarshaledStruct := new(Post)

	if err := json.Unmarshal(b, unmarshaledStruct); err != nil {
		return nil, err
	}

	return unmarshaledStruct, nil
}

package models

import (
	"encoding/json"
	"database/sql"
)

type User struct {
	About    string `json:"about,omitempty"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Nickname string `json:"nickname,omitempty"`
}

func UnmarshalUser(b []byte) (*User, error) {
	unmarshaledStruct := new(User)

	if err := json.Unmarshal(b, unmarshaledStruct); err != nil {
		return nil, err
	}

	return unmarshaledStruct, nil
}

func ScanUserFromRows(rows *sql.Rows) (*User, error) {
	user := new(User)
	err := rows.Scan(&user.About, &user.Email, &user.Fullname, &user.Nickname)

	return user, err
}

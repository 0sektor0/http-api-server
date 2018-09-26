package restapi

import (
	"database/sql"
	m "projects/http-api-server/models"

	_ "github.com/mattn/go-sqlite3"
)

type UsersStorage struct {
	db *sql.DB
}

func (s *UsersStorage) AddUser(nickname string, user *m.User) *ApiResponse { //(*m.User, []*m.User, *m.Error) {
	_, err := s.db.Query("INSERT INTO user (about, email, fullname) VALUES('?', '?', '?')", user.About, user.Email, user.Fullname)
	if err != nil {
		return &ApiResponse{Code: 500, Response: &m.Error{"500"}}
	}

	user.Nickname=nickname;
	return &ApiResponse{Code: 201, Response: user}
}

func (s *UsersStorage) UpdateUser(nickname string, update *m.UserUpdate) *ApiResponse { //(*m.User, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *UsersStorage) GetUserDetails(nickname string) *ApiResponse { //(*m.User, *m.Error) {
	panic("unemplimented function")
	return nil
}

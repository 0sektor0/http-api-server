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
	return &ApiResponse{Code: 200, Response: new(m.User)}
}

func (s *UsersStorage) UpdateUser(nickname string, update *m.UserUpdate) *ApiResponse { //(*m.User, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *UsersStorage) GetUserDetails(nickname string) *ApiResponse { //(*m.User, *m.Error) {
	panic("unemplimented function")
	return nil
}

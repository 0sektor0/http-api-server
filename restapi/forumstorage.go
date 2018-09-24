package restapi

import (
	"database/sql"
	m "projects/http-api-server/models"

	_ "github.com/mattn/go-sqlite3"
)

type ForumsStorage struct {
	db *sql.DB
}

func (s *ForumsStorage) AddForum(forum *m.Forum) *ApiResponse { //(*m.Forum, *m.Error) {}
	return &ApiResponse{Code: 200, Response: new(m.Forum)}
}

func (s *ForumsStorage) GetForumDetails(slug string) *ApiResponse { //(*m.Forum, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ForumsStorage) GetForumThreads(slug string, limit int, since string, desc bool) *ApiResponse { //([]*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ForumsStorage) GetForumUsers(slug string, limit int, since string, desc bool) *ApiResponse { //([]*m.User, *m.Error) {
	panic("unemplimented function")
	return nil
}

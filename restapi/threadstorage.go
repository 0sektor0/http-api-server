package restapi

import (
	"database/sql"
	m "projects/http-api-server/models"

	_ "github.com/mattn/go-sqlite3"
)

type ThreadsStorage struct {
	db *sql.DB
}

func (s *ThreadsStorage) GetThreadDetails(slug string) *ApiResponse { //(*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ThreadsStorage) GetThreadPosts(slug string, limit int, since int, sort string, desc bool) *ApiResponse { //([]*m.Post, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ThreadsStorage) UpdateThread(slug string, thread *m.ThreadUpdate) *ApiResponse { //(*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ThreadsStorage) VoteForThread(slug string, vote *m.Vote) *ApiResponse { //(*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ThreadsStorage) AddThread(slug string, thread *m.Thread) *ApiResponse { //*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

package main

import (
	"database/sql"
	m "projects/http-api-server/models"

	_ "github.com/mattn/go-sqlite3"
)

type ApiService struct {
	connector  string
	connection string
}

func NewApiService(connector string, connection string) *ApiService {
	service := &ApiService{
		connector:  connector,
		connection: connection,
	}

	return service
}

func (s *ApiService) OpenDbConnection() (*sql.DB, *ApiResponse) {
	db, err := sql.Open(s.connector, s.connection)
	if err != nil {
		return nil, &ApiResponse{
			Code:     500,
			Response: m.Error{err.Error()},
		}
	}

	return db, nil
}

func (s *ApiService) AddForum(forum *m.Forum) *ApiResponse { //(*m.Forum, *m.Error) {}
	db, err := s.OpenDbConnection()
	if(err != nil) {
		return err
	}
	defer db.Close()

	return &ApiResponse{ Code: 200, Response: new(m.Forum)}
}

func (s *ApiService) AddPosts(slug string, posts []m.Post) *ApiResponse { //([]*m.Post, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ApiService) AddThread(slug string, thread *m.Thread) *ApiResponse { //*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ApiService) AddUser(nickname string, user *m.User) *ApiResponse { //(*m.User, []*m.User, *m.Error) {
	db, err := s.OpenDbConnection()
	if(err != nil) {
		return err
	}
	defer db.Close()

	return &ApiResponse{ Code: 200, Response: new(m.User)}
}

func (s *ApiService) GetServiceStatus() *ApiResponse { //(*m.Status, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ApiService) GetForumDetails(slug string) *ApiResponse { //(*m.Forum, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ApiService) GetUserDetails(nickname string) *ApiResponse { //(*m.User, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ApiService) GetThreadDetails(slug string) *ApiResponse { //(*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ApiService) GetPostDetails(id int32, related []string) *ApiResponse { //(*m.PostFull, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ApiService) GetForumUsers(slug string, limit int, since string, desc bool) *ApiResponse { //([]*m.User, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ApiService) GetForumThreads(slug string, limit int, since string, desc bool) *ApiResponse { //([]*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ApiService) GetThreadPosts(slug string, limit int, since int, sort string, desc bool) *ApiResponse { //([]*m.Post, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ApiService) UpdatePost(id int64, update *m.PostUpdate) *ApiResponse { //(*m.Post, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ApiService) UpdateThread(slug string, thread *m.ThreadUpdate) *ApiResponse { //(*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ApiService) VipeServiceStatus() *ApiResponse { //*m.Error {
	panic("unemplimented function")
	return nil
}

func (s *ApiService) VoteForThread(slug string, vote *m.Vote) *ApiResponse { //(*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ApiService) UpdateUser(nickname string, update *m.UserUpdate) *ApiResponse { //(*m.User, *m.Error) {
	panic("unemplimented function")
	return nil
}

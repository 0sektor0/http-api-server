package api

import (
	"database/sql"
	m "projects/http-api-server/models"

	_ "github.com/mattn/go-sqlite3"
)

// ответ от сервиса работы с апи
type ApiResponse struct {
	Code     int
	Response interface{}
}

// интерфейс для взаимодействия с апи
type IApiService interface {
	AddForum(forum *m.Forum) *ApiResponse

	AddPosts(slug string, posts []m.Post) *ApiResponse

	AddThread(slug string, thread *m.Thread) *ApiResponse

	AddUser(nickname string, user *m.User) *ApiResponse

	GetServiceStatus() *ApiResponse

	GetForumDetails(slug string) *ApiResponse

	GetUserDetails(nickname string) *ApiResponse

	GetThreadDetails(slug string) *ApiResponse

	GetPostDetails(id int32, related []string) *ApiResponse

	GetForumUsers(slug string, limit int, since string, desc bool) *ApiResponse

	GetForumThreads(slug string, limit int, since string, desc bool) *ApiResponse

	GetThreadPosts(slug string, limit int, since int, sort string, desc bool) *ApiResponse

	UpdatePost(id int64, update *m.PostUpdate) *ApiResponse

	UpdateThread(slug string, thread *m.ThreadUpdate) *ApiResponse

	VipeServiceStatus() *ApiResponse

	VoteForThread(slug string, vote *m.Vote) *ApiResponse

	UpdateUser(nickname string, update *m.UserUpdate) *ApiResponse
}

// реализация интерфейса для работы с апи
// sql.DB не настоящий коннект к бд, а абстракция, которая управляет коннектами
type ApiService struct {
	db *sql.DB
}

func NewApiService(connector string, connection string) (*ApiService, error) {
	db, err := sql.Open(connector, connection)
	if err != nil {
		return nil, err
	}

	service := &ApiService{
		db: db,
	}

	return service, nil
}

func (s *ApiService) AddForum(forum *m.Forum) *ApiResponse { //(*m.Forum, *m.Error) {}
	return &ApiResponse{Code: 200, Response: new(m.Forum)}
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
	return &ApiResponse{Code: 200, Response: new(m.User)}
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

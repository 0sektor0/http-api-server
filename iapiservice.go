package main

import (
	m "projects/http-api-server/models"
)

type ApiResponse struct {
	Code     int
	Response interface{}
}

//this interface designed for interaction with api
type IApiService interface {
	AddForum(forum *m.Forum) *ApiResponse

	AddPosts(slug string, posts []*m.Post) *ApiResponse

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

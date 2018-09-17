package main

import (
	m "projects/http-api-server/models"
)

type ApiService struct {
}

func (service *ApiService) AddForum(forum *m.Forum) *ApiResponse { //(*m.Forum, *m.Error) {}
	panic("unemplimented function")
	return nil
}

func (service *ApiService) AddPosts(slug string, posts []*m.Post) *ApiResponse { //([]*m.Post, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (service *ApiService) AddThread(slug string, thread *m.Thread) *ApiResponse { //*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (service *ApiService) AddUser(nickname string, user *m.User) *ApiResponse { //(*m.User, []*m.User, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (service *ApiService) GetServiceStatus() *ApiResponse { //(*m.Status, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (service *ApiService) GetForumDetails(slug string) *ApiResponse { //(*m.Forum, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (service *ApiService) GetUserDetails(nickname string) *ApiResponse { //(*m.User, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (service *ApiService) GetThreadDetails(slug string) *ApiResponse { //(*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (service *ApiService) GetPostDetails(id int32, related []string) *ApiResponse { //(*m.PostFull, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (service *ApiService) GetForumUsers(slug string, limit int, since string, desc bool) *ApiResponse { //([]*m.User, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (service *ApiService) GetForumThreads(slug string, limit int, since string, desc bool) *ApiResponse { //([]*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (service *ApiService) GetThreadPosts(slug string, limit int, since int, sort string, desc bool) *ApiResponse { //([]*m.Post, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (service *ApiService) UpdatePost(id int64, update *m.PostUpdate) *ApiResponse { //(*m.Post, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (service *ApiService) UpdateThread(slug string, thread *m.ThreadUpdate) *ApiResponse { //(*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (service *ApiService) VipeServiceStatus() *ApiResponse { //*m.Error {
	panic("unemplimented function")
	return nil
}

func (service *ApiService) VoteForThread(slug string, vote *m.Vote) *ApiResponse { //(*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (service *ApiService) UpdateUser(nickname string, update *m.UserUpdate) *ApiResponse { //(*m.User, *m.Error) {
	panic("unemplimented function")
	return nil
}

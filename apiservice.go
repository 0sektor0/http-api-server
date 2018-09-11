package main

import (
	m "projects/http-api-server/models"
)

type ApiService struct {
}

func (service *ApiService) AddForum(forum *m.Forum) (*m.Forum, *m.Error) {
	panic("unemplimented function")
	return nil, nil
}

func (service *ApiService) AddPosts(slug string, posts []*m.Post) ([]*m.Post, *m.Error) {
	panic("unemplimented function")
	return nil, nil
}

func (service *ApiService) AddThread(slug string, thread *m.Thread) (*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil, nil
}

func (service *ApiService) AddUser(nickname string, user *m.User) (*m.User, []*m.User, *m.Error) {
	panic("unemplimented function")
	return nil, nil, nil
}

func (service *ApiService) GetServiceStatus() (*m.Status, *m.Error) {
	panic("unemplimented function")
	return nil, nil
}

func (service *ApiService) GetForumDetails(slug string) (*m.Forum, *m.Error) {
	panic("unemplimented function")
	return nil, nil
}

func (service *ApiService) GetUserDetails(nickname string) (*m.User, *m.Error) {
	panic("unemplimented function")
	return nil, nil
}

func (service *ApiService) GetThreadDetails(slug string) (*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil, nil
}

func (service *ApiService) GetPostDetails(id int32, related []string) (*m.PostFull, *m.Error) {
	panic("unemplimented function")
	return nil, nil
}

func (service *ApiService) GetForumUsers(slug string, limit int, since string, desc bool) ([]*m.User, *m.Error) {
	panic("unemplimented function")
	return nil, nil
}

func (service *ApiService) GetForumThreads(slug string, limit int, since string, desc bool) ([]*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil, nil
}

func (service *ApiService) GetThreadPosts(slug string, limit int, since int, sort string, desc bool) ([]*m.Post, *m.Error) {
	panic("unemplimented function")
	return nil, nil
}

func (service *ApiService) UpdatePost(id int64, update *m.PostUpdate) (*m.Post, *m.Error) {
	panic("unemplimented function")
	return nil, nil
}

func (service *ApiService) UpdateThread(slug string, thread *m.ThreadUpdate) (*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil, nil
}

func (service *ApiService) VipeServiceStatus() *m.Error {
	panic("unemplimented function")
	return nil
}

func (service *ApiService) VoteForThread(slug string, vote *m.Vote) (*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil, nil
}

func (service *ApiService) UpdateUser(nickname string, update *m.UserUpdate) (*m.User, *m.Error) {
	panic("unemplimented function")
	return nil, nil
}

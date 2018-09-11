package main

import (
	m "projects/http-api-server/models"
)

//this interface designed for interaction with api
type IApiService interface {
	AddForum(forum *m.Forum) (*m.Forum, *m.Error)

	AddPosts(slug string, posts []*m.Post) ([]*m.Post, *m.Error)

	AddThread(slug string, thread *m.Thread) (*m.Thread, *m.Error)

	AddUser(nickname string, user *m.User) (*m.User, []*m.User, *m.Error)

	GetServiceStatus() (*m.Status, *m.Error)

	GetForumDetails(slug string) (*m.Forum, *m.Error)

	GetUserDetails(nickname string) (*m.User, *m.Error)

	GetThreadDetails(slug string) (*m.Thread, *m.Error)

	GetPostDetails(id int32, related []string) (*m.PostFull, *m.Error)

	GetForumUsers(slug string, limit int, since string, desc bool) ([]*m.User, *m.Error)

	GetForumThreads(slug string, limit int, since string, desc bool) ([]*m.Thread, *m.Error)

	GetThreadPosts(slug string, limit int, since int, sort string, desc bool) ([]*m.Post, *m.Error)

	UpdatePost(id int64, update *m.PostUpdate) (*m.Post, *m.Error)

	UpdateThread(slug string, thread *m.ThreadUpdate) (*m.Thread, *m.Error)

	VipeServiceStatus() *m.Error

	VoteForThread(slug string, vote *m.Vote) (*m.Thread, *m.Error)

	UpdateUser(nickname string, update *m.UserUpdate) (*m.User, *m.Error)
}
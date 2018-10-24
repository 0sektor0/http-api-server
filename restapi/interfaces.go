package restapi

import (
	m "http-api-server/models"
)

type IUsersStorage interface {
	AddUser(user *m.User) *ApiResponse

	GetUserDetails(nickname string) *ApiResponse

	UpdateUser(nickname string, update *m.User) *ApiResponse

	GetUsersByForum(slug string, limit int, since string, desc bool) *ApiResponse
}

type IForumsStorage interface {
	AddForum(forum *m.Forum) *ApiResponse

	GetForumDetails(slug string) *ApiResponse
}

type IThreadsStorage interface {
	AddThread(thread *m.Thread) *ApiResponse

	GetThreadDetails(slug string) *ApiResponse

	GetThreadPosts(slug string, limit int, since int, sort string, desc bool) *ApiResponse

	UpdateThread(slug string, thread *m.Thread) *ApiResponse

	VoteForThread(slug string, vote *m.Vote) *ApiResponse

	GetThreadByForum(slug string, limit int, since string, desc bool) *ApiResponse
}

type IPostsStorage interface {
	AddPosts(slug string, posts []*m.Post) *ApiResponse

	GetPostDetails(id int, related []string) *ApiResponse

	UpdatePost(id int, update *m.PostUpdate) *ApiResponse
}

type IServiceStorege interface {
	GetServiceStatus() *ApiResponse

	VipeServiceStatus() *ApiResponse
}

package restapi

import (
	m "projects/http-api-server/models"
)

type IUsersStorage interface {
	AddUser(user *m.User) *ApiResponse

	GetUserDetails(nickname string) *ApiResponse

	UpdateUser(update *m.User) *ApiResponse
}

type IForumsStorage interface {
	AddForum(forum *m.Forum) *ApiResponse

	GetForumDetails(slug string) *ApiResponse

	GetForumUsers(slug string, limit int, since string, desc bool) *ApiResponse

	GetForumThreads(slug string, limit int, since string, desc bool) *ApiResponse
}

type IThreadsStorage interface {
	AddThread(slug string, thread *m.Thread) *ApiResponse

	GetThreadDetails(slug string) *ApiResponse

	GetThreadPosts(slug string, limit int, since int, sort string, desc bool) *ApiResponse

	UpdateThread(slug string, thread *m.ThreadUpdate) *ApiResponse

	VoteForThread(slug string, vote *m.Vote) *ApiResponse
}

type IPostsStorage interface {
	AddPosts(slug string, posts []m.Post) *ApiResponse

	GetPostDetails(id int32, related []string) *ApiResponse

	UpdatePost(id int64, update *m.PostUpdate) *ApiResponse
}

type IServiceStorege interface {
	GetServiceStatus() *ApiResponse

	VipeServiceStatus() *ApiResponse
}

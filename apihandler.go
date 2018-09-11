package main

import (
	"github.com/kataras/iris"
)

//this struct is proxy between network and api logic
type ApiHandler struct {
	apiService IApiService
}

func NewApiHandler() *ApiHandler {
	api := new(ApiHandler)
	api.apiService = new(ApiService)

	return api
}

func (handler *ApiHandler) AddForum(ctx iris.Context) {
	ctx.StatusCode(404)
	ctx.Write([]byte("404\n"))
}

func (handler *ApiHandler) AddPosts(ctx iris.Context) {
	ctx.StatusCode(404)
	ctx.Write([]byte("404\n"))
}

func (handler *ApiHandler) AddThread(ctx iris.Context) {
	ctx.StatusCode(404)
	ctx.Write([]byte("404\n"))
}

func (handler *ApiHandler) AddUser(ctx iris.Context) {
	ctx.StatusCode(404)
	ctx.Write([]byte("404\n"))
}

func (handler *ApiHandler) GetServiceStatus(ctx iris.Context) {
	ctx.StatusCode(404)
	ctx.Write([]byte("404\n"))
}

func (handler *ApiHandler) GetForumDetails(ctx iris.Context) {
	ctx.StatusCode(404)
	ctx.Write([]byte("404\n"))
}

func (handler *ApiHandler) GetUserDetails(ctx iris.Context) {
	ctx.StatusCode(404)
	ctx.Write([]byte("404\n"))
}

func (handler *ApiHandler) GetThreadDetails(ctx iris.Context) {
	ctx.StatusCode(404)
	ctx.Write([]byte("404\n"))
}

func (handler *ApiHandler) GetPostDetails(ctx iris.Context) {
	ctx.StatusCode(404)
	ctx.Write([]byte("404\n"))
}

func (handler *ApiHandler) GetForumUsers(ctx iris.Context) {
	ctx.StatusCode(404)
	ctx.Write([]byte("404\n"))
}

func (handler *ApiHandler) GetForumThreads(ctx iris.Context) {
	ctx.StatusCode(404)
	ctx.Write([]byte("404\n"))
}

func (handler *ApiHandler) GetThreadPosts(ctx iris.Context) {
	ctx.StatusCode(404)
	ctx.Write([]byte("404\n"))
}

func (handler *ApiHandler) UpdatePost(ctx iris.Context) {
	ctx.StatusCode(404)
	ctx.Write([]byte("404\n"))
}

func (handler *ApiHandler) UpdateThread(ctx iris.Context) {
	ctx.StatusCode(404)
	ctx.Write([]byte("404\n"))
}

func (handler *ApiHandler) VipeServiceStatus(ctx iris.Context) {
	ctx.StatusCode(404)
	ctx.Write([]byte("404\n"))
}

func (handler *ApiHandler) VoteForThread(ctx iris.Context) {
	ctx.StatusCode(404)
	ctx.Write([]byte("404\n"))
}

func (handler *ApiHandler) UpdateUser(ctx iris.Context) {
	ctx.StatusCode(404)
	ctx.Write([]byte("404\n"))
}

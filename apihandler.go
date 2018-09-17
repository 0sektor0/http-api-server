package main

import (
	"encoding/json"
	m "projects/http-api-server/models"

	"github.com/kataras/iris"
)

//this struct is proxy between network and api logic
type ApiHandler struct {
	apiService IApiService
}

func NewApiHandler() *ApiHandler {
	api := new(ApiHandler)
	api.apiService = NewApiService("sqlite3", "./data/forum.db")

	return api
}

func WriteResponse(response *ApiResponse, ctx iris.Context) {
	data, _ := json.Marshal(response.Response)

	ctx.StatusCode(response.Code)
	ctx.Write(data)
}

func (h *ApiHandler) AddForum(ctx iris.Context) {
	forum := new(m.Forum)
	ctx.ReadJSON(forum)

	WriteResponse(h.apiService.AddForum(forum), ctx)
}

func (h *ApiHandler) AddThread(ctx iris.Context) {
	slug := ctx.Params().Get("slug")
	thread := new(m.Thread)
	ctx.ReadJSON(thread)

	WriteResponse(h.apiService.AddThread(slug, thread), ctx)
}

func (h *ApiHandler) AddPosts(ctx iris.Context) {
	slug := ctx.Params().Get("slug_or_id")
	posts := new([]m.Post)
	ctx.ReadJSON(posts)

	WriteResponse(h.apiService.AddPosts(slug, *posts), ctx)
}

func (h *ApiHandler) AddUser(ctx iris.Context) {
	nickname := ctx.Params().Get("nickname")
	user := new(m.User)
	ctx.ReadJSON(user)

	WriteResponse(h.apiService.AddUser(nickname, user), ctx)
}

func (h *ApiHandler) GetServiceStatus(ctx iris.Context) {
	WriteResponse(h.apiService.GetServiceStatus(), ctx)
}

func (h *ApiHandler) GetForumDetails(ctx iris.Context) {
	slug := ctx.Params().Get("slug")
	WriteResponse(h.apiService.GetForumDetails(slug), ctx)
}

func (h *ApiHandler) GetUserDetails(ctx iris.Context) {
	nickname := ctx.Params().Get("nickname")
	WriteResponse(h.apiService.GetUserDetails(nickname), ctx)
}

func (h *ApiHandler) GetThreadDetails(ctx iris.Context) {
	slug := ctx.Params().Get("slug_or_id")
	limit, _ := ctx.URLParamInt("limit")
	since := ctx.URLParam("since")
	desc, _ := ctx.URLParamBool("desc")

	WriteResponse(h.apiService.GetForumThreads(slug, limit, since, desc), ctx)
}

func (h *ApiHandler) GetPostDetails(ctx iris.Context) {
	ctx.StatusCode(404)
	ctx.Write([]byte("404\n"))
}

func (h *ApiHandler) GetForumUsers(ctx iris.Context) {
	slug := ctx.Params().Get("slug")
	limit, _ := ctx.URLParamInt("limit")
	since := ctx.URLParam("since")
	desc, _ := ctx.URLParamBool("desc")

	WriteResponse(h.apiService.GetForumUsers(slug, limit, since, desc), ctx)
}

func (h *ApiHandler) GetForumThreads(ctx iris.Context) {
	slug := ctx.Params().Get("slug")
	limit, _ := ctx.URLParamInt("limit")
	since := ctx.URLParam("since")
	desc, _ := ctx.URLParamBool("desc")

	WriteResponse(h.apiService.GetForumThreads(slug, limit, since, desc), ctx)
}

func (h *ApiHandler) GetThreadPosts(ctx iris.Context) {
	slug := ctx.Params().Get("slug_or_id")
	limit, _ := ctx.URLParamInt("limit")
	since, _ := ctx.URLParamInt("since")
	sort := ctx.URLParam("sort")
	desc, _ := ctx.URLParamBool("desc")

	WriteResponse(h.apiService.GetThreadPosts(slug, limit, since, sort, desc), ctx)
}

func (h *ApiHandler) UpdatePost(ctx iris.Context) {
	id, _ := ctx.URLParamInt64("id")
	postUpdate := new(m.PostUpdate)

	WriteResponse(h.apiService.UpdatePost(id, postUpdate), ctx)
}

func (h *ApiHandler) UpdateThread(ctx iris.Context) {
	slug := ctx.URLParam("slug_or_id")
	threadUpdate := new(m.ThreadUpdate)

	WriteResponse(h.apiService.UpdateThread(slug, threadUpdate), ctx)
}

func (h *ApiHandler) VipeServiceStatus(ctx iris.Context) {
	WriteResponse(h.apiService.VipeServiceStatus(), ctx)
}

func (h *ApiHandler) VoteForThread(ctx iris.Context) {
	slug := ctx.URLParam("slug_or_id")
	vote := new(m.Vote)

	WriteResponse(h.apiService.VoteForThread(slug, vote), ctx)
}

func (h *ApiHandler) UpdateUser(ctx iris.Context) {
	nickname := ctx.URLParam("nickname")
	userUpdate := new(m.UserUpdate)

	WriteResponse(h.apiService.UpdateUser(nickname, userUpdate), ctx)
}

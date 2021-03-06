package main

import (
	"encoding/json"
	"github.com/kataras/iris"
	"log"
	"strconv"
	"strings"

	m "http-api-server/models"
	rest "http-api-server/restapi"
)

//посредник между сетью и логикой апи
type ApiHandler struct {
	apiService *rest.ApiService
}

func WriteResponse(response *rest.ApiResponse, ctx iris.Context) {
	data, err := json.Marshal(response.Response)
	if err != nil {
		log.Fatalln(err)
	}

	ctx.ContentType("application/json")
	ctx.StatusCode(response.Code)
	ctx.Write(data)
}

func (h *ApiHandler) AddForum(ctx iris.Context) {
	forum := new(m.Forum)
	ctx.ReadJSON(forum)

	WriteResponse(h.apiService.Forums.AddForum(forum), ctx)
}

func (h *ApiHandler) AddThread(ctx iris.Context) {
	thread := new(m.Thread)
	ctx.ReadJSON(thread)
	thread.Forum = ctx.Params().Get("slug")

	WriteResponse(h.apiService.Threads.AddThread(thread), ctx)
}

func (h *ApiHandler) AddPosts(ctx iris.Context) {
	slug := ctx.Params().Get("slug_or_id")
	posts := []*m.Post{}
	ctx.ReadJSON(&posts)

	WriteResponse(h.apiService.Posts.AddPosts(slug, posts), ctx)
}

func (h *ApiHandler) AddUser(ctx iris.Context) {
	nickname := ctx.Params().Get("nickname")
	user := new(m.User)

	ctx.ReadJSON(user)
	user.Nickname = nickname

	WriteResponse(h.apiService.Users.AddUser(user), ctx)
}

func (h *ApiHandler) GetServiceStatus(ctx iris.Context) {
	WriteResponse(h.apiService.Service.GetServiceStatus(), ctx)
}

func (h *ApiHandler) GetForumDetails(ctx iris.Context) {
	slug := ctx.Params().Get("slug")
	WriteResponse(h.apiService.Forums.GetForumDetails(slug), ctx)
}

func (h *ApiHandler) GetUserDetails(ctx iris.Context) {
	nickname := ctx.Params().Get("nickname")
	WriteResponse(h.apiService.Users.GetUserDetails(nickname), ctx)
}

func (h *ApiHandler) GetThreadDetails(ctx iris.Context) {
	slug := ctx.Params().Get("slug_or_id")
	WriteResponse(h.apiService.Threads.GetThreadDetails(slug), ctx)
}

func (h *ApiHandler) GetPostDetails(ctx iris.Context) {
	id, _ := strconv.Atoi(ctx.Params().Get("id"))
	rel := ctx.URLParam("related")

	WriteResponse(h.apiService.Posts.GetPostDetails(id, strings.Split(rel, ",")), ctx)
}

func (h *ApiHandler) GetForumUsers(ctx iris.Context) {
	slug := ctx.Params().Get("slug")
	limit, _ := ctx.URLParamInt("limit")
	since := ctx.URLParam("since")
	desc, _ := ctx.URLParamBool("desc")

	if limit < 0 {
		limit = 0
	}

	WriteResponse(h.apiService.Users.GetUsersByForum(slug, limit, since, desc), ctx)
}

func (h *ApiHandler) GetForumThreads(ctx iris.Context) {
	slug := ctx.Params().Get("slug")
	limit, _ := ctx.URLParamInt("limit")
	since := ctx.URLParam("since")
	desc, _ := ctx.URLParamBool("desc")

	if limit < 0 {
		limit = 0
	}

	WriteResponse(h.apiService.Threads.GetThreadByForum(slug, limit, since, desc), ctx)
}

func (h *ApiHandler) GetThreadPosts(ctx iris.Context) {
	slug := ctx.Params().Get("slug_or_id")
	limit, _ := ctx.URLParamInt("limit")
	since, _ := ctx.URLParamInt("since")
	sort := ctx.URLParam("sort")
	desc, _ := ctx.URLParamBool("desc")

	if limit < 0 {
		limit = 0
	}

	WriteResponse(h.apiService.Threads.GetThreadPosts(slug, limit, since, sort, desc), ctx)
}

func (h *ApiHandler) UpdatePost(ctx iris.Context) {
	id, _ := strconv.Atoi(ctx.Params().Get("id"))
	postUpdate := new(m.PostUpdate)
	ctx.ReadJSON(postUpdate)

	WriteResponse(h.apiService.Posts.UpdatePost(id, postUpdate), ctx)
}

func (h *ApiHandler) UpdateThread(ctx iris.Context) {
	thread := new(m.Thread)
	ctx.ReadJSON(thread)
	slug := ctx.Params().Get("slug_or_id")

	WriteResponse(h.apiService.Threads.UpdateThread(slug, thread), ctx)
}

func (h *ApiHandler) VipeServiceStatus(ctx iris.Context) {
	WriteResponse(h.apiService.Service.VipeServiceStatus(), ctx)
}

func (h *ApiHandler) VoteForThread(ctx iris.Context) {
	slug := ctx.Params().Get("slug_or_id")
	vote := new(m.Vote)
	ctx.ReadJSON(vote)

	WriteResponse(h.apiService.Threads.VoteForThread(slug, vote), ctx)
}

func (h *ApiHandler) UpdateUser(ctx iris.Context) {
	nickname := ctx.Params().Get("nickname")
	userUpdate := new(m.User)
	ctx.ReadJSON(userUpdate)

	WriteResponse(h.apiService.Users.UpdateUser(nickname, userUpdate), ctx)
}

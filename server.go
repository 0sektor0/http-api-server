package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func BuildServer() *iris.Application {
	api := NewApiHandler()

	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	app.Post("/forum/create", api.AddForum)
	app.Post("/forum/{slug:string}/create", api.AddThread)
	app.Get("/forum/{slug:string}/details", api.GetForumDetails)
	app.Get("/forum/{slug:string}/threads", api.GetForumThreads)
	app.Get("/forum/{slug:string}/users", api.GetForumUsers)
	app.Get("/post/{id:int}/details", api.GetPostDetails)
	app.Post("/post/{id}/details", api.UpdatePost)
	app.Post("/service/clear", api.VipeServiceStatus)
	app.Get("/service/status", api.GetServiceStatus)
	app.Post("/thread/{slug_or_id:string}/create", api.AddPosts)
	app.Get("/thread/{slug_or_id:string}/details", api.GetThreadDetails)
	app.Post("/thread/{slug_or_id:string}/details", api.UpdateThread)
	app.Get("/thread/{slug_or_id:string}/posts", api.GetThreadPosts)
	app.Post("/thread/{slug_or_id:string}/vote", api.VoteForThread)
	app.Post("/user/{nickname}/create", api.AddUser)
	app.Get("/user/{nickname}/profile", api.GetUserDetails)
	app.Post("/user/{nickname}/profile", api.UpdateUser)

	return app
}

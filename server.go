package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	api "projects/http-api-server/api"
)

func BuildServer(cfg *Configs) (*iris.Application, error) {
	apiService, err := api.NewApiService(cfg.Connector, cfg.Connection)
	if err != nil {
		return nil, err
	}

	api := &ApiHandler {
		apiService: apiService,
	}

	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	app.Post("api/forum/create", api.AddForum)
	app.Post("api/forum/{slug:string}/create", api.AddThread)
	app.Get("api/forum/{slug:string}/details", api.GetForumDetails)
	app.Get("api/forum/{slug:string}/threads", api.GetForumThreads)
	app.Get("api/forum/{slug:string}/users", api.GetForumUsers)
	app.Get("api/post/{id:int}/details", api.GetPostDetails)
	app.Post("api/post/{id}/details", api.UpdatePost)
	app.Post("api/service/clear", api.VipeServiceStatus)
	app.Get("api/service/status", api.GetServiceStatus)
	app.Post("api/thread/{slug_or_id:string}/create", api.AddPosts)
	app.Get("api/thread/{slug_or_id:string}/details", api.GetThreadDetails)
	app.Post("api/thread/{slug_or_id:string}/details", api.UpdateThread)
	app.Get("api/thread/{slug_or_id:string}/posts", api.GetThreadPosts)
	app.Post("api/thread/{slug_or_id:string}/vote", api.VoteForThread)
	app.Post("api/user/{nickname}/create", api.AddUser)
	app.Get("api/user/{nickname}/profile", api.GetUserDetails)
	app.Post("api/user/{nickname}/profile", api.UpdateUser)

	return app, nil
}

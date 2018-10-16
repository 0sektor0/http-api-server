package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	rest "projects/http-api-server/restapi"
)

func BuildServer(cfg *Configs) (*iris.Application, error) {
	apiService, err := rest.NewApiService(cfg.Connector, cfg.Connection)
	if err != nil {
		return nil, err
	}

	apiHandler := &ApiHandler{
		apiService: apiService,
	}

	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	//success
	app.Post("api/user/{nickname}/create", apiHandler.AddUser)
	app.Get("api/user/{nickname}/profile", apiHandler.GetUserDetails)
	app.Post("api/user/{nickname}/profile", apiHandler.UpdateUser)
	app.Post("api/forum/create", apiHandler.AddForum)
	app.Get("api/forum/{slug:string}/details", apiHandler.GetForumDetails)
	app.Post("api/forum/{slug:string}/create", apiHandler.AddThread)
	app.Get("api/forum/{slug:string}/threads", apiHandler.GetForumThreads)
	app.Post("api/thread/{slug_or_id:string}/create", apiHandler.AddPosts)
	app.Get("api/forum/{slug:string}/users", apiHandler.GetForumUsers)
	app.Get("api/post/{id:int}/details", apiHandler.GetPostDetails)
	app.Post("api/post/{id}/details", apiHandler.UpdatePost)
	app.Post("api/service/clear", apiHandler.VipeServiceStatus)
	app.Get("api/service/status", apiHandler.GetServiceStatus)
	app.Get("api/thread/{slug_or_id:string}/details", apiHandler.GetThreadDetails)
	app.Post("api/thread/{slug_or_id:string}/details", apiHandler.UpdateThread)
	app.Get("api/thread/{slug_or_id:string}/posts", apiHandler.GetThreadPosts)
	app.Post("api/thread/{slug_or_id:string}/vote", apiHandler.VoteForThread)

	return app, nil
}

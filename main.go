package main

import (
	"log"

	"github.com/kataras/iris"
)

const cfgPath = "./data/cfg.json"

func main() {
	cfg, err := LoadConfigs(cfgPath)
	if err != nil {
		cfg = &Configs{
			Connector: "postgres",
			Connection: "host=127.0.0.1 port=5432 user=forum_admin password=forum_admin dbname=forum sslmode=disable",
			Port: ":5000",
		}
	}

	app, err := BuildServer(cfg)
	if err != nil {
		log.Fatalf("failed to build server: %s", err)
		return
	}

	log.Println("START")
	app.Run(iris.Addr(cfg.Port), iris.WithoutServerError(iris.ErrServerClosed))
}

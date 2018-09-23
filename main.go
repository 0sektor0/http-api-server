package main

import (
	"log"

	"github.com/kataras/iris"
)

const cfgPath = "./data/cfg.json"

func main() {
	cfg, err := LoadConfigs(cfgPath)
	if err != nil {
		log.Fatalf("failed to read configuration file: %s", err)
		return
	}

	app, err := BuildServer(cfg)
	if err != nil {
		log.Fatalf("failed to build server: %s", err)
		return
	}

	app.Run(iris.Addr(cfg.Port), iris.WithoutServerError(iris.ErrServerClosed))
}

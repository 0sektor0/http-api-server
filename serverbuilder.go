package main

import (
	"github.com/kataras/iris"
)

type IServerBuilder interface {
	Build(cfgFile string) *iris.Application
}
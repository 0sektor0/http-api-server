package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := BuildServer()
	app.Run(iris.Addr(":5000"), iris.WithoutServerError(iris.ErrServerClosed))
}

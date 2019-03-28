package main

import (
	"fmt"
	"time"

	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	app.StaticWeb("/static", "./web")

	tmpl := iris.HTML("./web/templates", ".html").Reload(true)
	app.RegisterView(tmpl)

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.ViewData("YEAR", time.Now().Year())

		ctx.View("index.html")
	})

	app.Handle("GET", "/q/{topic}", func(ctx iris.Context) {
		app.Logger().Info(fmt.Sprintf("Questions about topics %+v", ctx.Params()))
		ctx.HTML(fmt.Sprintf("<h1>Questions about %s </h1>", ctx.Params().Get("topic")))
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

}

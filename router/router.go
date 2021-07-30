package router

import "github.com/kataras/iris/v12"

// SetRouter -
func SetRouter(app *iris.Application) {
	r := app.Party("/api")

	setCookie(r)
}

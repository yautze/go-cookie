package router

import (
	"go-cookie/controller/cookie"
	"go-cookie/middle"

	"github.com/kataras/iris/v12"
)

func setCookie(e iris.Party) {
	m := e.Party("/cookie")
	{
		m.Get("/set", middle.HandleFunc(cookie.SetCookie))
		m.Get("/read", middle.HandleFunc(cookie.ReadCookie))
	}
}

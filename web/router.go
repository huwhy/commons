package web

import (
	"github.com/huwhy/commons/basemodel"
	"github.com/kataras/iris/v12"
)

type JsonHandler func(ctx iris.Context) *basemodel.Json

type Router struct {
	iris.Party
}

func NewRouter(app iris.Party) *Router {
	return &Router{app}
}

func (r *Router) GET(path string, handler JsonHandler) {
	r.Get(path, JsonHandle(handler))
}

func (r *Router) POST(path string, handler JsonHandler) {
	r.Post(path, JsonHandle(handler))
}

func (r *Router) DELETE(path string, handler JsonHandler) {
	r.Delete(path, JsonHandle(handler))
}

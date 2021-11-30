package web

import (
	"github.com/huwhy/commons/basemodel"
	"github.com/kataras/iris/v12"
)

func JsonHandle(fn func(c iris.Context) *basemodel.Json, interceptors ...Interceptor) func(c iris.Context) {
	return func(c iris.Context) {
		if len(interceptors) > 0 {
			for _, interceptor := range interceptors {
				if ok, msg := interceptor(c); !ok {
					c.JSON(basemodel.JsonFail(msg))
					return
				}
			}
		}
		json := fn(c)
		c.JSON(json)
	}
}

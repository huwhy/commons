package web

import (
	"github.com/kataras/iris/v12"
	"huwhy.cn/commons/basemodel"
)

func JsonHandle(fn func(c iris.Context) *basemodel.Json) func(c iris.Context) {
	return func(c iris.Context) {
		json := fn(c)
		c.JSON(json)
	}
}

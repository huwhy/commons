package web

import (
	"git.huwhy.cn/commons/basemodel"
	"github.com/kataras/iris/v12"
)

func JsonHandle(fn func(c iris.Context) *basemodel.Json) func(c iris.Context) {
	return func(c iris.Context) {
		json := fn(c)
		c.JSON(json)
	}
}

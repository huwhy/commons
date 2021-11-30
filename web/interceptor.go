package web

import (
	"github.com/kataras/iris/v12"
)

type Interceptor func(c iris.Context) (bool, string)

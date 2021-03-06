package web

import (
	"github.com/huwhy/commons/basemodel"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"runtime"
)

func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}

func ErrorFilter(log *zap.SugaredLogger) func(iris.Context) {
	return func(context iris.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("[%d] err={%+v}, %s", GetWebTraceId(context), err, GetErrStack())
				context.JSON(basemodel.JsonFail("系统异常，请联系服务人员"))
			}
		}()
		context.Next()
	}
}

func GetErrStack() string {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	return string(buf[:n])
}

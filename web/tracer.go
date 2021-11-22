package web

import (
	"context"
	"github.com/huwhy/commons/constant"
	"github.com/kataras/iris/v12"
	"time"
)

func GetWebTraceId(ctx iris.Context) int64 {
	return GetTraceId(GetTracer(ctx))
}

func GetTraceId(ctx context.Context) int64 {
	if ctx == nil {
		return 0
	}
	if v, ok := ctx.Value(constant.TRACER_ID).(int64); ok {
		return v
	} else {
		return 0
	}
}

func SetTracer(ctx iris.Context) context.Context {
	newContext := NewTracerContext()
	ctx.Values().Set(constant.TRACER_CTX, newContext)
	return newContext
}

func GetTracer(ctx iris.Context) context.Context {
	if c, ok := ctx.Values().Get(constant.TRACER_CTX).(context.Context); ok {
		return c
	} else {
		return SetTracer(ctx)
	}
}

func NewTracerContext() context.Context {
	return context.WithValue(context.Background(), constant.TRACER_ID, time.Now().UnixNano())
}

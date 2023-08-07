package sentry

import (
	"log"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

type Sentry struct {
	Dsn         string `protobuf:"bytes,1,opt,name=dsn,proto3" json:"dsn,omitempty"`
	Environment string `protobuf:"bytes,2,opt,name=environment,proto3" json:"environment,omitempty"`
}

func SentryInit(SentryCfg *Sentry) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              SentryCfg.Dsn,
		AttachStacktrace: true,
		Environment:      SentryCfg.Environment,
	})
	if err != nil {
		log.Println("启动sentry失败,请配置正确参数", err.Error())
	}
}

func Middleware(rePanic bool) gin.HandlerFunc {
	// 这里是指sentry上报逻辑，所以要开启true，因为这里截胡后，后面的中间件就不知道有没有panic了
	return sentrygin.New(sentrygin.Options{
		Repanic: rePanic,
	})
}

func MiddlewareTag(ctx *gin.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			hub.Scope().SetTag("someRandomTag", "maybeYouNeedIt")
		}
		ctx.Next()
	}
}

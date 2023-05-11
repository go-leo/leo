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

func Middleware() gin.HandlerFunc {
	return sentrygin.New(sentrygin.Options{
		Repanic: true,
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

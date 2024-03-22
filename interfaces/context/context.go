package interfaces_context

import (
	"github.com/aamirmousavi/dong/internal/database/mongodb"
	"github.com/gin-gonic/gin"
)

const CONTEXT = "app_context"

type Context interface {
	Mongo() *mongodb.Handler
}

func GetAppContext(ctx *gin.Context) Context {
	return ctx.MustGet(CONTEXT).(Context)
}

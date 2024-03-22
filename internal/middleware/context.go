package middleware

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	"github.com/aamirmousavi/dong/internal/context"
	"github.com/gin-gonic/gin"
)

type contextHandler struct {
	ctx *context.Context
}

func NewContextHandler(ctx *context.Context) *contextHandler {
	return &contextHandler{
		ctx,
	}
}

func (ch *contextHandler) AppContext() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(interfaces_context.CONTEXT, ch.ctx)

		ctx.Next()
	}
}

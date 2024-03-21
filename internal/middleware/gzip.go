package middleware

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Gzip() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		gzip.Gzip(gzip.DefaultCompression)(ctx)
	}
}

package router

import (
	api_user "github.com/aamirmousavi/dong/internal/api/user"
	"github.com/aamirmousavi/dong/internal/database/mongodb"
	"github.com/aamirmousavi/dong/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Run(
	mongodb *mongodb.Handler,
	addr string,
) error {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Recovery())

	router.Use(middleware.CORS())

	router.Use(middleware.Gzip())

	api_user.Configure(
		router.Group("/api/user"),
	)

	return router.Run(addr)
}

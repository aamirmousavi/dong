package router

import (
	api_user "github.com/aamirmousavi/dong/internal/api/user"
	"github.com/aamirmousavi/dong/internal/context"
	"github.com/aamirmousavi/dong/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Run(
	appContext *context.Context,
	addr string,
) error {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Recovery())

	router.Use(middleware.CORS())

	router.Use(middleware.Gzip())

	router.Use(middleware.NewContextHandler(appContext).AppContext())

	authorizationMiddleware := middleware.NewAuthorizationHandler(appContext).Authorization

	api_user.Configure(
		authorizationMiddleware,
		router.Group("/api/user"),
	)

	return router.Run(addr)
}

package router

import (
	"os"

	api_contact "github.com/aamirmousavi/dong/internal/api/contact"
	api_period "github.com/aamirmousavi/dong/internal/api/period"
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

	router.Use(middleware.AllowOptions())

	api := router.Group("/api")
	api.Use(middleware.CORS())

	api.Use(middleware.Gzip())

	api.Use(middleware.NewContextHandler(appContext).AppContext())

	authorizationMiddleware := middleware.NewAuthorizationHandler(appContext).Authorization

	{
		api_user.Configure(
			authorizationMiddleware,
			api.Group("/user"),
		)

		api.Use(authorizationMiddleware())

		api_contact.Configure(
			api.Group("/contact"),
		)

		api_period.Configure(
			api.Group("/period"),
		)
	}

	storage := router.Group("/storage")

	storage.GET("/*path", func(ctx *gin.Context) {
		data, err := os.ReadFile(
			"/storage" + ctx.Param("path"),
		)
		if err != nil {
			ctx.JSON(404, gin.H{
				"message": "فایل مورد نظر یافت نشد",
			})
			return
		}
		ctx.Writer.Write(data)
	})

	return router.Run(addr)
}

package user

import (
	"github.com/gin-gonic/gin"
)

func Configure(
	authorizationMiddleware func() gin.HandlerFunc,
	group *gin.RouterGroup,
) {
	group.POST("/register", register)
	group.POST("/register_authorization", registerAuthorization)
	group.POST("/login", login)
	group.POST("/login_authorization", loginAuthorization)
	group.GET("/user_exists", userExists)

	group.Use(authorizationMiddleware())

	group.GET("/profile", profile)
	group.POST("/logout", logout)
}

package contacts

import "github.com/gin-gonic/gin"

func Configure(
	group *gin.RouterGroup,
) {
	group.POST("/add", add)
	group.GET("/list", list)
	group.DELETE("/remove", remove)
	group.GET("/get", get)
}

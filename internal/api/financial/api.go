package financial

import "github.com/gin-gonic/gin"

func Configure(
	group *gin.RouterGroup,
) {
	group.GET("/get", get)
}

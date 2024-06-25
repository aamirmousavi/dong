package payment

import (
	"github.com/gin-gonic/gin"
)

func Configure(
	group *gin.RouterGroup,
) {
	group.GET("/list", list)
	group.POST("/add", add)
	group.GET("/get", get)
	group.POST("/update", update)
}

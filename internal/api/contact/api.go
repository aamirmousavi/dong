package contact

import "github.com/gin-gonic/gin"

func Configure(group *gin.RouterGroup) {
	group.POST("/add", add)
	group.PUT("/edit", edit)
	group.DELETE("/remove", remove)
	group.GET("/list", list)
	group.GET("/get", get)
}

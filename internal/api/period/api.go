package period

import (
	api_factor "github.com/aamirmousavi/dong/internal/api/period/factor"
	api_payment "github.com/aamirmousavi/dong/internal/api/period/payment"
	"github.com/gin-gonic/gin"
)

func Configure(
	group *gin.RouterGroup,
) {
	group.GET("/list", list)
	group.POST("/add", add)
	group.GET("/get", get)
	group.PUT("/add_user", addUser)

	api_factor.Configure(
		group.Group("/factor"),
	)

	api_payment.Configure(
		group.Group("/payment"),
	)

}

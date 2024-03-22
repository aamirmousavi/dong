package user

import (
	interfaces_profile "github.com/aamirmousavi/dong/interfaces/profile"
	"github.com/gin-gonic/gin"
)

func profile(ctx *gin.Context) {
	profile := interfaces_profile.GetProfile(ctx)

	ctx.JSON(200, gin.H{
		"message": "با موفقیت درخواست شد",
		"data":    profile,
	})
}

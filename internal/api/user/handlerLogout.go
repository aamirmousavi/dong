package user

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	interfaces_profile "github.com/aamirmousavi/dong/interfaces/profile"
	"github.com/gin-gonic/gin"
)

func logout(ctx *gin.Context) {
	profile := interfaces_profile.GetProfile(ctx)
	app := interfaces_context.GetAppContext(ctx)

	if err := app.Mongo().UserHandler.Logout(ctx, profile.AccessToken); err != nil {
		ctx.JSON(500, gin.H{
			"message": "خطای داخلی",
			"desc":    err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "با موفقیت خارج شدید",
	})
}

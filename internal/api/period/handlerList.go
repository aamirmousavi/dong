package period

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	interfaces_profile "github.com/aamirmousavi/dong/interfaces/profile"
	"github.com/gin-gonic/gin"
)

func list(ctx *gin.Context) {
	app := interfaces_context.GetAppContext(ctx)
	profile := interfaces_profile.GetProfile(ctx)

	list, err := app.Mongo().PeroidHandler.GetByUserId(profile.User.Id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"data": list,
	})
}

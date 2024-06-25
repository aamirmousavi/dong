package user

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	interfaces_profile "github.com/aamirmousavi/dong/interfaces/profile"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func bankGet(ctx *gin.Context) {
	app := interfaces_context.GetAppContext(ctx)
	profile := interfaces_profile.GetProfile(ctx)
	bank, err := app.Mongo().UserHandler.GetBank(profile.User.Id)
	if err != nil && err != mongo.ErrNoDocuments {
		ctx.JSON(500, gin.H{
			"message": "خطای داخلی",
			"desc":    err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "با موفقیت درخواست شد",
		"data":    bank,
	})
}

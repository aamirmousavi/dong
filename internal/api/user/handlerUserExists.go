package user

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/gin-gonic/gin"
)

type userExistsRequest struct {
	Number string `form:"number" binding:"required"`
}

func userExists(ctx *gin.Context) {
	params, err := bind.Bind[userExistsRequest](ctx)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "ورودی نامعتبر است",
			"desc":    err.Error(),
		})
		return
	}
	app := interfaces_context.GetAppContext(ctx)
	exist, err := app.Mongo().UserHandler.UserExists(ctx, params.Number)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "خطای داخلی",
			"desc":    err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"exist": exist,
	})
}

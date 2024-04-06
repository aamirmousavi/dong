package contact

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	interfaces_profile "github.com/aamirmousavi/dong/interfaces/profile"
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/gin-gonic/gin"
)

type listRequest struct {
	Limit  *int64 `form:"limit"`
	Offest *int64 `form:"offest"`
}

func list(ctx *gin.Context) {
	p, err := bind.Bind[listRequest](ctx)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "ورودی نامعتبر است",
			"desc":    err.Error(),
		})
		return
	}
	var limit int64 = 100
	var offest int64
	if p.Limit != nil && *p.Limit != 0 {
		limit = *p.Limit
	}
	app := interfaces_context.GetAppContext(ctx)
	profile := interfaces_profile.GetProfile(ctx)
	cnts, err := app.Mongo().ContactHandler.List(
		ctx,
		profile.User.Id,
		limit,
		offest,
	)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "خطای داخلی",
			"desc":    err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "لیست مخاطبین با موفقیت دریافت شد",
		"data":    cnts,
	})
}

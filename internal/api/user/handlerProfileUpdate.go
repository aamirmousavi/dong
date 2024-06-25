package user

import (
	"mime/multipart"

	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	interfaces_profile "github.com/aamirmousavi/dong/interfaces/profile"
	"github.com/aamirmousavi/dong/service/file/image"
	"github.com/gin-gonic/gin"
)

type profileUpdateRequest struct {
	FirstName string                `form:"first_name" binding:"required"`
	LastName  string                `form:"last_name" binding:"required"`
	Pic       *multipart.FileHeader `form:"pic"`
}

func profileUpdate(ctx *gin.Context) {
	params := new(profileUpdateRequest)
	if err := ctx.Bind(params); err != nil {
		ctx.JSON(400, gin.H{
			"message": "ورودی نامعتبر است",
			"desc":    err.Error(),
		})
		return
	}
	app := interfaces_context.GetAppContext(ctx)
	profile := interfaces_profile.GetProfile(ctx)
	usr := profile.User
	usr.FirstName = params.FirstName
	usr.LastName = params.LastName
	if params.Pic != nil {
		iamgeAddr, err := image.Profile(params.Pic, usr.Id.Hex())
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "خطای داخلی",
				"desc":    err.Error(),
			})
			return
		}
		usr.Pic = &iamgeAddr
	}
	if err := app.Mongo().UserHandler.Update(ctx, usr); err != nil {
		ctx.JSON(500, gin.H{
			"message": "خطای داخلی",
			"desc":    err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "با موفقیت انجام شد",
	})
}

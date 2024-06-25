package user

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type loginAuthorizationRequest struct {
	Code   int    `form:"code" binding:"required"`
	Number string `form:"number" binding:"required"`
}

func loginAuthorization(ctx *gin.Context) {
	params, err := bind.Bind[loginAuthorizationRequest](ctx)
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
	if !exist {
		ctx.JSON(400, gin.H{
			"message": "کاربر با این شماره ثبت نام نکرده است",
		})
		return
	}
	if err := app.Mongo().OTPHandler.Check(ctx, params.Number, params.Code, nil); err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(400, gin.H{
				"message": "کد وارد شده اشتباه است",
			})
			return
		}
		ctx.JSON(500, gin.H{
			"message": "خطای داخلی",
			"desc":    err.Error(),
		})
		return
	}
	usr, err := app.Mongo().UserHandler.Get(ctx, params.Number)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "خطای داخلی",
			"desc":    err.Error(),
		})
		return
	}
	tkn := usr.GenerateToken()
	if err := app.Mongo().UserHandler.CreateToken(ctx, tkn); err != nil {
		ctx.JSON(500, gin.H{
			"message": "خطای داخلی",
			"desc":    err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "ثبت نام با موفقیت انجام شد",
		"data":    tkn,
	})
}

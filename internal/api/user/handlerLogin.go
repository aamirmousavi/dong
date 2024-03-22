package user

import (
	"fmt"

	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	"github.com/aamirmousavi/dong/internal/database/mongodb/otp"
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/aamirmousavi/dong/utils/rand"
	"github.com/aamirmousavi/dong/utils/sms"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Number string `form:"number" binding:"required"`
}

func login(ctx *gin.Context) {
	params, err := bind.Bind[loginRequest](ctx)
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
	code := rand.IntStandard()
	otp := otp.New(
		params.Number,
		code,
	)
	if err := app.Mongo().OTPHandler.Create(ctx, otp); err != nil {
		ctx.JSON(500, gin.H{
			"message": "خطای داخلی",
			"desc":    err.Error(),
		})
		return
	}
	if err := sms.Send(
		params.Number,
		fmt.Sprintf(
			`به دونگ خوش آمدید
کد فعال سازی شما: %d
`,
			code,
		),
	); err != nil {
		ctx.JSON(500, gin.H{
			"message": "خطای داخلی",
			"desc":    err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "کد فعال سازی با موفقیت ارسال شد",
	})
}

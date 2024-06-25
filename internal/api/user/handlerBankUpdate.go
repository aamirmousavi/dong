package user

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	interfaces_profile "github.com/aamirmousavi/dong/interfaces/profile"
	"github.com/aamirmousavi/dong/internal/database/mongodb/user"
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/gin-gonic/gin"
)

type bankUpdateRequest struct {
	CardNumber  string `form:"card_number" binding:"required"`
	BankName    string `form:"bank_name" binding:"required"`
	AccountName string `form:"account_name" binding:"required"`
	ShebaNumber string `form:"sheba_number" binding:"required"`
}

func bankUpdate(ctx *gin.Context) {
	params, err := bind.Bind[bankUpdateRequest](ctx)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "ورودی نامعتبر است",
			"desc":    err.Error(),
		})
		return
	}
	app := interfaces_context.GetAppContext(ctx)
	profile := interfaces_profile.GetProfile(ctx)
	bank := user.NewBank(
		profile.User.Id,
		params.CardNumber,
		params.BankName,
		params.AccountName,
		params.ShebaNumber,
	)
	if err := app.Mongo().UserHandler.UpdateBank(bank); err != nil {
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

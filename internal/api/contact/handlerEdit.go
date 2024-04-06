package contact

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	interfaces_profile "github.com/aamirmousavi/dong/interfaces/profile"
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type editRequest struct {
	Id        string  `json:"id" binding:"required"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Number    *string `json:"number"`
}

func edit(ctx *gin.Context) {
	p, err := bind.BindJson[editRequest](ctx)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "ورودی نامعتبر است",
			"desc":    err.Error(),
		})
		return
	}
	if (p.FirstName == nil || *p.FirstName == "") && (p.LastName == nil || *p.LastName == "") && (p.Number == nil || *p.Number == "") {
		ctx.JSON(400, gin.H{
			"message": "یکی از سه فیلد نام خوانوادگی و شماره باید مقدار داشته باشد",
		})
		return
	}
	oid, err := primitive.ObjectIDFromHex(p.Id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "شناسه مخاطب معتبر نیست",
		})
		return
	}
	app := interfaces_context.GetAppContext(ctx)
	profile := interfaces_profile.GetProfile(ctx)
	cnt, err := app.Mongo().ContactHandler.Get(
		ctx,
		oid,
		profile.User.Id,
	)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "خطای داخلی",
			"desc":    err.Error(),
		})
		return
	}
	if (cnt.Number == nil || *cnt.Number == "") && (p.Number != nil || *p.Number != "") {
		contactUserId, err := app.Mongo().GetId(ctx, *p.Number)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "خطای داخلی",
				"desc":    err.Error(),
			})
			return
		}
		cnt.ContactUserId = contactUserId
	}
	cnt.FirstName = p.FirstName
	cnt.LastName = p.LastName
	if err := app.Mongo().ContactHandler.Edit(ctx, cnt); err != nil {
		ctx.JSON(500, gin.H{
			"message": "خطای داخلی",
			"desc":    err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "مخاطب با موفقیت ویرایش شد",
	})
}

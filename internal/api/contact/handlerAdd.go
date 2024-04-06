package contact

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	interfaces_profile "github.com/aamirmousavi/dong/interfaces/profile"
	"github.com/aamirmousavi/dong/internal/database/mongodb/contact"
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type addRequest struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Number    *string `json:"number"`
}

func add(ctx *gin.Context) {
	p, err := bind.BindJson[addRequest](ctx)
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
	app := interfaces_context.GetAppContext(ctx)
	profile := interfaces_profile.GetProfile(ctx)
	var contactUserId *primitive.ObjectID
	if p.Number != nil {
		contactUserId, err = app.Mongo().GetId(ctx, *p.Number)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "خطای داخلی",
				"desc":    err.Error(),
			})
			return
		}
	}
	cnt := contact.New(
		profile.User.Id,
		contactUserId,
		p.FirstName,
		p.LastName,
		p.Number,
		nil,
	).GenerateId()
	if err := app.Mongo().ContactHandler.Create(ctx, cnt); err != nil {
		ctx.JSON(500, gin.H{
			"message": "خطای داخلی",
			"desc":    err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "مخاطب با موفقیت ساخته شد",
	})
}

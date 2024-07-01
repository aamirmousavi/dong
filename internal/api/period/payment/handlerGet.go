package payment

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	"github.com/aamirmousavi/dong/lib"
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type getId struct {
	Id string `form:"id" binding:"required"`
}

func get(ctx *gin.Context) {
	p, err := bind.Bind[getId](ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	app := interfaces_context.GetAppContext(ctx)
	id, err := primitive.ObjectIDFromHex(p.Id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	payment, err := app.Mongo().BalanceHandler.PaymentGet(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sourceUser, err := app.Mongo().UserHandler.GetById(ctx, payment.SourceUserId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	targetUser, err := app.Mongo().UserHandler.GetById(ctx, payment.TargetUserId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	payment.SourceUserName = lib.Ptr(sourceUser.FirstName + " " + sourceUser.LastName)
	payment.TargetUserName = lib.Ptr(targetUser.FirstName + " " + targetUser.LastName)
	ctx.JSON(200, gin.H{"payment": payment})
}

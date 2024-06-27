package payment

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type listRequest struct {
	PeroidId *string `form:"peroid_id"`
	UserId   *string `form:"user_id"`
}

func list(ctx *gin.Context) {
	p, err := bind.Bind[listRequest](ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if p.PeroidId == nil && p.UserId == nil {
		ctx.JSON(400, gin.H{"error": "peroid_id or user_id is required"})
		return
	}
	app := interfaces_context.GetAppContext(ctx)
	var peroidId, userId *primitive.ObjectID
	if p.PeroidId != nil {
		pid, err := primitive.ObjectIDFromHex(*p.PeroidId)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		peroidId = &pid
	}
	if p.UserId != nil {
		uid, err := primitive.ObjectIDFromHex(*p.UserId)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		userId = &uid
	}
	payments, err := app.Mongo().BalanceHandler.PaymentList(peroidId, userId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": payments})
}

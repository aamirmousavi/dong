package contacts

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	interfaces_profile "github.com/aamirmousavi/dong/interfaces/profile"
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type getRequest struct {
	Id string `form:"id" binding:"required"`
}

func get(ctx *gin.Context) {
	p, err := bind.Bind[getRequest](ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	app := interfaces_context.GetAppContext(ctx)
	oid, err := primitive.ObjectIDFromHex(p.Id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	cnct, err := app.Mongo().ContactHandler.Get(oid)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	profile := interfaces_profile.GetProfile(ctx)
	peroids, err := app.Mongo().PeroidHandler.ListWithContact(
		profile.User.Id,
		cnct.UserId,
	)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	payments, err := app.Mongo().BalanceHandler.PaymentListWithContact(
		profile.User.Id,
		cnct.UserId,
	)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"data":     cnct,
		"peroids":  peroids,
		"payments": payments,
	})
}

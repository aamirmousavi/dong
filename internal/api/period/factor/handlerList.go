package factor

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	"github.com/aamirmousavi/dong/lib"
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type listRequest struct {
	PeroidId string `form:"peroid_id" binding:"required"`
}

func list(ctx *gin.Context) {
	p, err := bind.Bind[listRequest](ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	app := interfaces_context.GetAppContext(ctx)
	oid, err := primitive.ObjectIDFromHex(p.PeroidId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	factors, err := app.Mongo().FactorList(oid)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	for _, f := range factors {
		buyerUser, err := app.Mongo().UserHandler.GetById(ctx, f.Buyer)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		f.BuyerName = lib.Ptr(buyerUser.FirstName + " " + buyerUser.LastName)
	}
	ctx.JSON(200, gin.H{
		"data": factors,
	})
}

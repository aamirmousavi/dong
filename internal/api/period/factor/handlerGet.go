package factor

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	"github.com/aamirmousavi/dong/lib"
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

	factor, err := app.Mongo().PeroidHandler.FactorGet(oid)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	buyerUser, err := app.Mongo().UserHandler.GetById(ctx, factor.Buyer)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	factor.BuyerName = lib.Ptr(buyerUser.FirstName + " " + buyerUser.LastName)
	ctx.JSON(200, gin.H{
		"data": factor,
	})
}

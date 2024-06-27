package period

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userAddRequest struct {
	PeroidId string `form:"peroid_id" binding:"required"`
	UserId   string `form:"user_id" binding:"required"`
}

func userAdd(ctx *gin.Context) {
	p, err := bind.Bind[userAddRequest](ctx)
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
	userId, err := primitive.ObjectIDFromHex(p.UserId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := app.Mongo().PeroidHandler.AddUser(oid, userId); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"data": "ok",
	})
}

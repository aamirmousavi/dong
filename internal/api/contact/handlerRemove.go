package contacts

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type removeRequest struct {
	Id string `form:"id" binding:"required"`
}

func remove(ctx *gin.Context) {
	p, err := bind.Bind[removeRequest](ctx)
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
	if err := app.Mongo().ContactHandler.Remove(oid); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "با موفقیت حذف شد",
	})
}

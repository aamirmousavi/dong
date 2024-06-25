package period

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type addUserRequest struct {
	Id     string `json:"id" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}

func addUser(ctx *gin.Context) {
	p, err := bind.BindJson[addUserRequest](ctx)
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

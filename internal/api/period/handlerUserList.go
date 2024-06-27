package period

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userListRequest struct {
	PeroidId string `form:"peroid_id" binding:"required"`
}

func userList(ctx *gin.Context) {
	p, err := bind.Bind[userListRequest](ctx)
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
	peroid, err := app.Mongo().PeroidHandler.Get(oid)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	users, err := app.Mongo().UserHandler.GetMany(ctx, peroid.UserIds)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	type Person struct {
		Id         string  `json:"id"`
		Name       string  `json:"name"`
		MoneySpend uint64  `json:"money_spend"`
		Demand     *uint64 `json:"demand"`
		Debt       *uint64 `json:"debt"`
	}

	var periodUsers []Person
	for _, user := range users {
		periodUsers = append(periodUsers, Person{
			Id:         user.Id.Hex(),
			Name:       user.FirstName + " " + user.LastName,
			MoneySpend: 1000,
			Demand:     nil,
			Debt:       nil,
		})
	}

	ctx.JSON(200, gin.H{
		"data": periodUsers,
	})
}

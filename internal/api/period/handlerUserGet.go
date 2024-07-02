package period

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userGetRequest struct {
	PeroidId string `form:"peroid_id" binding:"required"`
	UserId   string `form:"user_id" binding:"required"`
}

func userGet(ctx *gin.Context) {
	app := interfaces_context.GetAppContext(ctx)
	p, err := bind.Bind[userGetRequest](ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	peroidId, err := primitive.ObjectIDFromHex(p.PeroidId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userId, err := primitive.ObjectIDFromHex(p.UserId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userBalance, err := app.Mongo().PeroidHandler.FactorCalculatedBalanceGetByUser(peroidId, userId)
	if err != nil && err != mongo.ErrNoDocuments {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	type PersonStat struct {
		Id     string `json:"id"`
		Title  string `json:"title"`
		Demand *int   `json:"demand"`
		Debt   *int   `json:"debt"`
	}
	personStat := make([]PersonStat, 0)
	if userBalance != nil && userBalance.ReletiveFactorCalculatedBalances != nil {
		for _, balance := range *userBalance.ReletiveFactorCalculatedBalances {
			user, err := app.Mongo().UserHandler.GetById(ctx, balance.UserId)
			if err != nil {
				ctx.JSON(500, gin.H{"error": err.Error()})
				return
			}
			personStat = append(personStat, PersonStat{
				Id:     balance.UserId.Hex(),
				Title:  user.FirstName + " " + user.LastName,
				Demand: balance.Demand,
				Debt:   balance.Debt,
			})
		}
	}
	ctx.JSON(200, gin.H{
		"data": personStat,
	})
}

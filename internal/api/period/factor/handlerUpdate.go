package factor

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	interfaces_profile "github.com/aamirmousavi/dong/interfaces/profile"
	"github.com/aamirmousavi/dong/internal/database/mongodb/balance"
	"github.com/aamirmousavi/dong/internal/database/mongodb/peroid"
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type updateRequest struct {
	Id    string `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required"`
	Buyer string `json:"buyer" binding:"required"`
	Users []struct {
		UserId      string `json:"user_id" binding:"required"`
		Coefficient int    `json:"coefficient" binding:"required"`
	} `json:"users" binding:"required"`
}

func update(ctx *gin.Context) {
	p, err := bind.BindJson[updateRequest](ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	app := interfaces_context.GetAppContext(ctx)
	profile := interfaces_profile.GetProfile(ctx)
	id, err := primitive.ObjectIDFromHex(p.Id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	UserWithCoefficient := make([]peroid.UserWithCoefficient, len(p.Users))
	for i, user := range p.Users {
		oid, err := primitive.ObjectIDFromHex(user.UserId)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		UserWithCoefficient[i] = peroid.UserWithCoefficient{
			UserId:      oid,
			Coefficient: user.Coefficient,
		}
	}
	buyer, err := primitive.ObjectIDFromHex(p.Buyer)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	oldFactor, err := app.Mongo().FactorGet(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	factor := peroid.NewFactor(
		p.Title,
		profile.User.Id,
		p.Price,
		buyer,
		UserWithCoefficient,
		oldFactor.PeroidId,
	).SetId(id)

	peroidData, err := app.Mongo().PeroidHandler.GetWithFactors(oldFactor.PeroidId, nil)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	peroidData.EditFactor(factor, oldFactor, false)

	if err := app.Mongo().PeroidHandler.FactorCalculatedBalanceAdd(&peroidData.Id, peroidData.Balances); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	balanceList := make(balance.BalanceList, 0)
	for _, calBalacne := range *peroidData.Balances {
		if calBalacne.Demand != nil {
			balanceList = append(balanceList, balance.NewBalance(
				&peroidData.Id,
				calBalacne.UserId,
				calBalacne.UserId,
				*calBalacne.Demand,
				false,
			))
			if calBalacne.ReletiveFactorCalculatedBalances != nil {
				for _, reletiveCalBalance := range *calBalacne.ReletiveFactorCalculatedBalances {
					if reletiveCalBalance.Debt != nil {
						balanceList = append(balanceList, balance.NewBalance(
							&peroidData.Id,
							calBalacne.UserId,
							reletiveCalBalance.UserId,
							*reletiveCalBalance.Debt,
							false,
						))
					}
				}
			}
		}
	}
	if err := app.Mongo().BalanceHandler.Add(&peroidData.Id, balanceList); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := app.Mongo().FactorUpdate(factor); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": factor})
}

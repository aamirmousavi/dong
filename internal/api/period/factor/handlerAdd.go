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

type addRequest struct {
	PeroidId string `json:"peroid_id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Price    int    `json:"price" binding:"required"`
	Buyer    string `json:"buyer" binding:"required"`
	Users    []struct {
		UserId      string `json:"user_id" binding:"required"`
		Coefficient int    `json:"coefficient" binding:"required"`
	} `json:"users" binding:"required"`
}

func add(ctx *gin.Context) {
	p, err := bind.BindJson[addRequest](ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	app := interfaces_context.GetAppContext(ctx)
	profile := interfaces_profile.GetProfile(ctx)

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
	peroidId, err := primitive.ObjectIDFromHex(p.PeroidId)
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
		peroidId,
	).GenerateId()

	peroid, err := app.Mongo().PeroidHandler.GetWithFactors(peroidId, nil)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	payments, err := app.Mongo().BalanceHandler.PaymentList(&peroidId, nil)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	peroid.Payments = &payments
	peroid.AddFactor(factor)
	if err := app.Mongo().PeroidHandler.FactorCalculatedBalanceAdd(&peroidId, peroid.Balances); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	balanceList := make(balance.BalanceList, 0)
	for _, calBalacne := range *peroid.Balances {
		if calBalacne.Demand != nil {
			balanceList = append(balanceList, balance.NewBalance(
				&peroidId,
				calBalacne.UserId,
				calBalacne.UserId,
				*calBalacne.Demand,
				false,
			))
			for _, reletiveCalBalance := range *calBalacne.ReletiveFactorCalculatedBalances {
				balanceList = append(balanceList, balance.NewBalance(
					&peroidId,
					calBalacne.UserId,
					reletiveCalBalance.UserId,
					*reletiveCalBalance.Debt,
					false,
				))
			}
		}
	}

	if err := app.Mongo().BalanceHandler.Add(&peroidId, balanceList); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	{
		peroid.Factors = nil
		peroid.Balances = nil
	}

	if err := app.Mongo().PeroidHandler.UpdateAll(peroid); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := app.Mongo().PeroidHandler.FactorAdd(factor); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, factor)
}

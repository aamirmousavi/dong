package payment

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
	Title        string  `form:"title"`
	PeroidId     *string `form:"peroid_id"`
	SourceUserId *string `form:"source_user_id"`
	TargetUserId string  `form:"target_user_id"`
	Amount       int     `form:"amount"`
}

func add(ctx *gin.Context) {
	p, err := bind.Bind[addRequest](ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	app := interfaces_context.GetAppContext(ctx)
	var PeroidId *primitive.ObjectID
	var sourceUserId primitive.ObjectID
	if p.PeroidId == nil {
		profile := interfaces_profile.GetProfile(ctx)
		sourceUserId = profile.User.Id
	} else {
		pid, err := primitive.ObjectIDFromHex(*p.PeroidId)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		PeroidId = &pid
	}
	if p.SourceUserId != nil {
		sourceUserId, err = primitive.ObjectIDFromHex(*p.SourceUserId)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}
	targetUserId, err := primitive.ObjectIDFromHex(p.TargetUserId)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	payment := balance.NewPayment(
		p.Title,
		PeroidId,
		sourceUserId,
		targetUserId,
		p.Amount,
	).GenerateId()
	if err := app.Mongo().BalanceHandler.PaymentAdd(payment); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if payment.PeroidId != nil {
		peroidData, err := app.Mongo().PeroidHandler.GetWithFactors(*payment.PeroidId, nil)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		peroidData.AddFactor(&peroid.Factor{
			Id:    payment.Id,
			Price: payment.Amount,
			Buyer: payment.SourceUserId,
			Users: []peroid.UserWithCoefficient{
				{
					UserId:      payment.TargetUserId,
					Coefficient: 1,
				},
			},
			PeroidId: *payment.PeroidId,
		}, true)
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
	}

	ctx.JSON(200, gin.H{"data": payment})

}

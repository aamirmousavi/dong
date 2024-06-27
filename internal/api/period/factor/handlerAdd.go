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
	Price    uint64 `json:"price" binding:"required"`
	Buyer    string `json:"buyer" binding:"required"`
	Users    []struct {
		UserId      string `json:"user_id" binding:"required"`
		Coefficient uint64 `json:"coefficient" binding:"required"`
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

	/*
		Go create balance for each user in factor
		user.each(
			userPrice = (p.price / sumofCoefficient) * user.coefficient
		)
	*/

	balanceList := make(balance.BalanceList, 0)
	sumOfCoefficient := uint64(0)
	for _, user := range factor.Users {
		sumOfCoefficient += user.Coefficient
	}
	for _, user := range factor.Users {
		userPrice := (factor.Price / sumOfCoefficient) * user.Coefficient
		balanceList = append(balanceList, balance.NewBalance(
			factor.Buyer,
			user.UserId,
			userPrice,
		))
	}

	if err := app.Mongo().BalanceHandler.Add(balanceList...); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := app.Mongo().PeroidHandler.FactorAdd(factor); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, factor)

	/*
		Reaclculate money
		user [A] must give user [B] 100$
		and user [B] must give user [C] 20$
		and user [C] must give user [A] 10$

		we remove bepending depths
		then user [A] must give user [B] 100$ - 10$ = 90$
		and user [B] must give user [C] 20$ - 10$ = 10$

		factor.each(

		)
	*/

}

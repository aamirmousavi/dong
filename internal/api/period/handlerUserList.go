package period

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	interfaces_profile "github.com/aamirmousavi/dong/interfaces/profile"
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
		MoneySpend int     `json:"money_spend"`
		Demand     *int    `json:"demand"`
		Debt       *int    `json:"debt"`
		CardNumber *string `json:"card_number"`
	}

	balances, err := app.Mongo().PeroidHandler.FactorCalculatedBalanceGet(oid)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	profile := interfaces_profile.GetProfile(ctx)
	profileCard, err := app.Mongo().GetBank(profile.User.Id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	profilePerson := Person{
		Id:         profile.User.Id.Hex(),
		Name:       profile.User.FirstName + " " + profile.User.LastName,
		MoneySpend: peroid.MoneySpend[profile.User.Id],
		Demand:     nil,
		Debt:       nil,
		CardNumber: &profileCard.CardNumber,
	}
	profileBalance, ok := balances.Find(profile.User.Id)
	if ok {
		if profileBalance.Demand != nil {
			profilePerson.Demand = profileBalance.Demand
		}
		if profileBalance.Debt != nil {
			profilePerson.Debt = profileBalance.Debt
		}
	}
	periodUsers := []Person{profilePerson}
	for _, user := range users {
		card, err := app.Mongo().GetBank(user.Id)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		userPerson := Person{
			Id:         user.Id.Hex(),
			Name:       user.FirstName + " " + user.LastName,
			MoneySpend: peroid.MoneySpend[user.Id],
			Demand:     nil,
			Debt:       nil,
			CardNumber: &card.CardNumber,
		}
		balance, ok := balances.Find(user.Id)
		if ok {
			if balance.Demand != nil {
				userPerson.Demand = balance.Demand
			}
			if balance.Debt != nil {
				userPerson.Debt = balance.Debt
			}
		}

		periodUsers = append(periodUsers, userPerson)
	}

	ctx.JSON(200, gin.H{
		"data": periodUsers,
	})
}

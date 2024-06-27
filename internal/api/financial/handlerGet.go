package financial

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	interfaces_profile "github.com/aamirmousavi/dong/interfaces/profile"
	"github.com/gin-gonic/gin"
)

func get(ctx *gin.Context) {
	app := interfaces_context.GetAppContext(ctx)
	profile := interfaces_profile.GetProfile(ctx)
	payments, err := app.Mongo().BalanceHandler.PaymentList(
		nil,
		&profile.User.Id,
	)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	factors, err := app.Mongo().PeroidHandler.FactorListByUser(profile.User.Id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"total_demand": 1_000,
		"total_debt":   2_000,
		"payments":     payments,
		"factors":      factors,
	})

}

package payment

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	interfaces_profile "github.com/aamirmousavi/dong/interfaces/profile"
	"github.com/aamirmousavi/dong/internal/database/mongodb/balance"
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type updateRequest struct {
	Id           string  `form:"id" binding:"required"`
	Title        string  `form:"title"`
	PeroidId     *string `form:"peroid_id"`
	SourceUserId *string `form:"source_user_id"`
	TargetUserId string  `form:"target_user_id"`
	Amount       uint64  `form:"amount"`
}

func update(ctx *gin.Context) {
	p, err := bind.Bind[updateRequest](ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	app := interfaces_context.GetAppContext(ctx)
	id, err := primitive.ObjectIDFromHex(p.Id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var PeroidId *primitive.ObjectID
	var sourceUserId primitive.ObjectID
	if p.PeroidId != nil {
		pid, err := primitive.ObjectIDFromHex(*p.PeroidId)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		PeroidId = &pid
	} else {
		profile := interfaces_profile.GetProfile(ctx)
		sourceUserId = profile.User.Id
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
	).SetId(id)

	if err := app.Mongo().BalanceHandler.PaymentUpdate(payment); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": payment})

}

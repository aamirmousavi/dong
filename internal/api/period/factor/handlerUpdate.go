package factor

import (
	"github.com/gin-gonic/gin"
)

type updateRequest struct {
	Id    string `json:"id" binding:"required"`
	Title string `json:"title" binding:"required"`
	Price uint64 `json:"price" binding:"required"`
	Buyer string `json:"buyer" binding:"required"`
	Users []struct {
		UserId      string `json:"user_id" binding:"required"`
		Coefficient uint64 `json:"coefficient" binding:"required"`
	} `json:"users" binding:"required"`
}

func update(ctx *gin.Context) {
	// p, err := bind.BindJson[updateRequest](ctx)
	// if err != nil {
	// 	ctx.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }
	// app := interfaces_context.GetAppContext(ctx)
	// profile := interfaces_profile.GetProfile(ctx)

	// UserWithCoefficient := make([]peroid.UserWithCoefficient, len(p.Users))
	// for i, user := range p.Users {
	// 	oid, err := primitive.ObjectIDFromHex(user.UserId)
	// 	if err != nil {
	// 		ctx.JSON(400, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	UserWithCoefficient[i] = peroid.UserWithCoefficient{
	// 		UserId:      oid,
	// 		Coefficient: user.Coefficient,
	// 	}
	// }
}

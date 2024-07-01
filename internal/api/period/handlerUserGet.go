package period

import (
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/gin-gonic/gin"
)

type userGetRequest struct {
	PeroidId string `form:"peroid_id" binding:"required"`
	UserID   string `form:"user_id" binding:"required"`
}

func userGet(ctx *gin.Context) {
	_, err := bind.Bind[userGetRequest](ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	type PersonStat struct {
		Id     string  `json:"id"`
		Title  string  `json:"title"`
		Demand *uint64 `json:"demand"`
		Debt   *uint64 `json:"debt"`
	}
	ctx.JSON(200, gin.H{
		"data": []PersonStat{
			{
				Id:     "5f7b7b7b7b7b7b7b7b7b7b7b",
				Title:  "Aamir Mousavi",
				Demand: ptr(1000),
				Debt:   nil,
			},
			{
				Id:     "5f7b7b7b7b7b7b7b7b7b7b7c",
				Title:  "Milad Rezaei",
				Demand: ptr(2000),
				Debt:   nil,
			},
			{
				Id:     "5f7b7b7b7b7b7b7b7b7b7b7A",
				Title:  "Milad Rezaei",
				Demand: nil,
				Debt:   ptr(4000),
			},
		},
	})
}

func ptr(v uint64) *uint64 {
	return &v
}

package period

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	interfaces_profile "github.com/aamirmousavi/dong/interfaces/profile"
	"github.com/aamirmousavi/dong/internal/database/mongodb/peroid"
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type addRequest struct {
	Title   string   `form:"title" binding:"required"`
	UserIds []string `form:"user_ids" binding:"required"`
}

func add(ctx *gin.Context) {
	p, err := bind.Bind[addRequest](ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	oids := make([]primitive.ObjectID, len(p.UserIds))
	for i, id := range p.UserIds {
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		oids[i] = oid
	}
	app := interfaces_context.GetAppContext(ctx)
	profile := interfaces_profile.GetProfile(ctx)
	contacts, err := app.Mongo().ContactHandler.GetByIds(oids)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	userIds := make([]primitive.ObjectID, 0)
	for _, c := range contacts {
		userIds = append(userIds, c.UserId)
	}
	per := peroid.NewPeroid(
		profile.User.Id,
		p.Title,
		userIds,
	).GenerateId()
	per.UserCount = uint64(len(oids)) + 1
	if err := app.Mongo().PeroidHandler.Add(per); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"data": per,
	})
}

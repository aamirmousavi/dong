package middleware

import (
	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	interfaces_profile "github.com/aamirmousavi/dong/interfaces/profile"
	"github.com/aamirmousavi/dong/internal/context"
	"github.com/aamirmousavi/dong/utils/bind"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type authorizationHandler struct {
	ctx *context.Context
}

func NewAuthorizationHandler(ctx *context.Context) *authorizationHandler {
	return &authorizationHandler{
		ctx,
	}
}

type authorizationRequest struct {
	Authorization string `header:"Authorization" binding:"required"`
}

func (ah *authorizationHandler) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params, err := bind.BindHeader[authorizationRequest](ctx)
		if err != nil {
			ctx.JSON(400, gin.H{
				"message": "ورودی نامعتبر است",
				"desc":    err.Error(),
			})
			ctx.Abort()
			return
		}
		app := interfaces_context.GetAppContext(ctx)

		usr, err := app.Mongo().UserHandler.GetUserByToken(ctx, params.Authorization)
		if err != nil && err != mongo.ErrNoDocuments {
			ctx.JSON(500, gin.H{
				"message": "خطای داخلی",
				"desc":    err.Error(),
			})
			ctx.Abort()
			return
		}

		if err == mongo.ErrNoDocuments {
			ctx.JSON(401, gin.H{
				"message": "توکن شما معتبر نیست",
			})
			ctx.Abort()
			return
		}

		ctx.Set(interfaces_profile.PROFILE, interfaces_profile.New(
			usr,
			params.Authorization,
		))

		ctx.Next()
	}
}

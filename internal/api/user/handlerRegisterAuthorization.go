package user

import (
	"mime/multipart"

	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	"github.com/aamirmousavi/dong/internal/database/mongodb/user"
	"github.com/aamirmousavi/dong/service/file/image"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type registerAuthorizationRequest struct {
	Code      int                   `form:"code" binding:"required"`
	Number    string                `form:"number" binding:"required"`
	FirstName string                `form:"first_name" binding:"required"`
	LastName  string                `form:"last_name" binding:"required"`
	Pic       *multipart.FileHeader `form:"pic"`
}

func registerAuthorization(ctx *gin.Context) {
	params := new(registerAuthorizationRequest)
	if err := ctx.Bind(params); err != nil {
		ctx.JSON(400, gin.H{
			"message": "ورودی نامعتبر است",
			"desc":    err.Error(),
		})
		return
	}
	app := interfaces_context.GetAppContext(ctx)
	exist, err := app.Mongo().UserHandler.UserExists(ctx, params.Number)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "خطای داخلی",
			"desc":    err.Error(),
		})
		return
	}
	if exist {
		ctx.JSON(400, gin.H{
			"message": "کاربر با این شماره قبلا ثبت نام کرده است",
		})
		return
	}
	if err := app.Mongo().OTPHandler.Check(ctx, params.Number, params.Code, nil); err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(400, gin.H{
				"message": "کد وارد شده اشتباه است",
			})
			return
		}
		ctx.JSON(500, gin.H{
			"message": "خطای داخلی",
			"desc":    err.Error(),
		})
		return
	}
	id := primitive.NewObjectID()
	var pic *string
	if params.Pic != nil {
		iamgeAddr, err := image.Profile(params.Pic, id.Hex())
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "خطای داخلی",
				"desc":    err.Error(),
			})
			return
		}
		pic = &iamgeAddr
	}
	usr := user.NewUser(params.FirstName, params.LastName, params.Number, pic).SetId(id)
	if err := app.Mongo().UserHandler.Create(ctx, usr); err != nil {
		ctx.JSON(500, gin.H{
			"message": "خطای داخلی",
			"desc":    err.Error(),
		})
		return
	}
	tkn := usr.GenerateToken()
	if err := app.Mongo().UserHandler.CreateToken(ctx, tkn); err != nil {
		ctx.JSON(500, gin.H{
			"message": "خطای داخلی",
			"desc":    err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "ثبت نام با موفقیت انجام شد",
		"data":    tkn,
	})
}

package contacts

import (
	"mime/multipart"

	interfaces_context "github.com/aamirmousavi/dong/interfaces/context"
	interfaces_profile "github.com/aamirmousavi/dong/interfaces/profile"
	"github.com/aamirmousavi/dong/internal/database/mongodb/contact"
	"github.com/aamirmousavi/dong/internal/database/mongodb/user"
	"github.com/aamirmousavi/dong/service/file/image"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type addRequest struct {
	Number    string                `form:"number" binding:"required"`
	FirstName *string               `form:"first_name"`
	LastName  *string               `form:"last_name"`
	Pic       *multipart.FileHeader `form:"pic"`
}

func add(ctx *gin.Context) {
	params := new(addRequest)
	if err := ctx.Bind(params); err != nil {
		ctx.JSON(400, gin.H{
			"message": "ورودی نامعتبر است",
			"desc":    err.Error(),
		})
		return
	}
	app := interfaces_context.GetAppContext(ctx)
	profile := interfaces_profile.GetProfile(ctx)

	exist, err := app.Mongo().UserHandler.UserExists(ctx, params.Number)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "خطای داخلی",
			"desc":    err.Error(),
		})
		return
	}
	var cnct *contact.Contact
	if exist {
		contactUser, err := app.Mongo().UserHandler.Get(ctx, params.Number)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": "خطای داخلی",
				"desc":    err.Error(),
			})
			return
		}
		cnct = contact.NewContact(
			params.Number,
			contactUser.FirstName,
			contactUser.LastName,
			contactUser.Pic,
			contactUser.Id,
		).GenerateId()
	} else {
		id := primitive.NewObjectID()
		var pic *string
		if params.Pic != nil {
			imageAddr, err := image.Profile(params.Pic, id.Hex())
			if err != nil {
				ctx.JSON(500, gin.H{
					"message": "خطای داخلی",
					"desc":    err.Error(),
				})
				return
			}
			pic = &imageAddr
		}
		newUser := user.NewUser(params.Number, *params.FirstName, *params.LastName, pic).SetId(id)
		if err := app.Mongo().UserHandler.Create(ctx, newUser); err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		cnct = contact.NewContact(
			params.Number,
			*params.FirstName,
			*params.LastName,
			pic,
			id,
		).GenerateId()
	}

	cnct.ContactOf = profile.User.Id
	if err := app.Mongo().ContactHandler.Add(cnct); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "با موفقیت اضافه شد",
		"data":    cnct,
	})
}

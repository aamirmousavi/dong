package profile

import (
	"github.com/aamirmousavi/dong/internal/database/mongodb/user"
	"github.com/gin-gonic/gin"
)

const PROFILE = "profile"

type Profile struct {
	User        *user.User
	AccessToken string
}

func New(
	usr *user.User,
	accessToken string,
) *Profile {
	return &Profile{
		User:        usr,
		AccessToken: accessToken,
	}
}

func GetProfile(ctx *gin.Context) Profile {
	return ctx.MustGet(PROFILE).(Profile)
}
